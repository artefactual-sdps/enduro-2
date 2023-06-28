package db

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"time"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/artefactual-sdps/enduro/internal/storage/persistence"
)

// Wait up to five minutes is another process is already on it.
const lockTimeout = time.Minute * 5

// Connect returns a database handler.
func Connect(driver, dsn string) (db *sql.DB, err error) {
	switch driver {
	case "mysql":
		db, err = ConnectMySQL(dsn)
	case "sqlite3":
		db, err = ConnectSQLite(dsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %q", driver)
	}

	return db, err
}

// ConnectMySQL returns a MySQL database handler which is safe for concurrent
// access.
func ConnectMySQL(dsn string) (db *sql.DB, err error) {
	config, err := mysqldriver.ParseDSN(dsn)
	if err != nil {
		return nil, fmt.Errorf("error parsing dsn: %w (%s)", err, dsn)
	}
	config.Collation = "utf8mb4_unicode_ci"
	config.Loc = time.UTC
	config.ParseTime = true
	config.MultiStatements = true
	config.Params = map[string]string{
		"time_zone": "UTC",
	}

	conn, err := mysqldriver.NewConnector(config)
	if err != nil {
		return nil, fmt.Errorf("error creating connector: %w", err)
	}

	db = sql.OpenDB(conn)

	// Set reasonable defaults in the built-in pool.
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	// Register Prometheus collector.
	c := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: "src",
			Subsystem: "mysql_app_" + config.DBName,
			Name:      "open_connections",
			Help:      "Number of open connections to MySQL DB, as reported by mysql.DB.Stats()",
		},
		func() float64 {
			s := db.Stats()
			return float64(s.OpenConnections)
		},
	)
	prometheus.MustRegister(c)

	return db, nil
}

// ConnectSQLite returns a SQLlite database handler which is NOT safe for
// concurrent access.
func ConnectSQLite(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", dsn)

	conns := runtime.NumCPU()
	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns)
	db.SetConnMaxLifetime(0)
	db.SetConnMaxIdleTime(0)

	pragmas := []string{
		"journa_mode=WAL",
		"synchronous=OFF",
		"foreign_keys=ON",
		"tempo_store=MEMORY",
		"busy_timeout=1000", // Used with "_txlock=immediate" or "BEGIN IMMEDIATE".
	}
	for _, pragma := range pragmas {
		if _, err := db.Exec("PRAGMA " + pragma); err != nil {
			return nil, err
		}
	}

	return db, err
}

func MigrateEnduroDatabase(db *sql.DB) error {
	src, err := enduroSourceDriver()
	if err != nil {
		return fmt.Errorf("error loading migration sources: %v", err)
	}

	return up(db, src)
}

func MigrateEnduroStorageDatabase(db *sql.DB) error {
	src, err := persistence.SourceDriver()
	if err != nil {
		return fmt.Errorf("error loading migration sources: %v", err)
	}

	return up(db, src)
}

// up migrates the database.
func up(db *sql.DB, src source.Driver) error {
	_, ok := db.Driver().(*mysqldriver.MySQLDriver)
	if !ok {
		return fmt.Errorf("only MySQL migrations are supported")
	}

	m, err := newMigrate(db, src)
	if err != nil {
		return fmt.Errorf("error creating golang-migrate object: %v", err)
	}

	err = doMigrate(m)
	if err != nil {
		return fmt.Errorf("error during database migration: %v", err)
	}

	return nil
}

// newMigrate builds the golang-migrate object.
func newMigrate(db *sql.DB, src source.Driver) (*migrate.Migrate, error) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("error creating migrate driver: %w", err)
	}

	m, err := migrate.NewWithInstance("source", src, "driver", driver)
	if err != nil {
		return nil, fmt.Errorf("error creating migrate instance: %w", err)
	}

	m.LockTimeout = lockTimeout

	return m, nil
}

func doMigrate(m *migrate.Migrate) error {
	err := m.Up()
	if err == nil || err == migrate.ErrNoChange {
		return nil
	}

	if os.IsNotExist(err) {
		_, dirty, verr := m.Version()
		if verr != nil {
			return verr
		}
		if dirty {
			return err
		}
		return nil
	}

	return err
}
