package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const (
	// defaultSize is the default number of packages to generate.
	defaultSize = 10000

	// batchSize is the number of rows to insert with a single database query.
	batchSize = 1000
)

type database interface {
	connect() error
	init() error
	batchInsertPkgs([]pkg) (int, error)
	insertPkgs([]pkg) (int, error)
}

type Config struct{}

func main() {
	args := os.Args

	n := defaultSize
	if len(args) > 2 {
		fail("Error: too many arguments")
	}
	if len(args) > 1 {
		i, err := strconv.Atoi(args[1])
		if err != nil {
			fail("Error: first argument must be a number")
		}
		n = i
	}

	db := &mysqlDB{
		cfg: mysql.Config{
			User:      "enduro",
			Passwd:    "enduro123",
			Net:       "tcp",
			Addr:      "127.0.0.1:3306",
			DBName:    "enduro",
			ParseTime: true,
		},
	}
	c := newCmd(Config{}, db)

	out, err := c.run(n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func usage() {
	fmt.Printf(`
Usage: go run main.go [COUNT]
Generates COUNT (default: %d) sample packages for testing.
`,
		defaultSize,
	)
}

func fail(msg string) {
	usage()
	fmt.Println(msg)
	os.Exit(1)
}

type pkg struct {
	ID          string
	Name        string
	WorkflowID  string
	RunID       uuid.UUID
	AIPID       uuid.UUID
	LocationID  uuid.UUID
	Status      string
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
}

func (p pkg) toSlice() []any {
	return []any{
		p.ID,
		p.Name,
		p.WorkflowID,
		p.RunID,
		p.AIPID,
		p.LocationID,
		p.Status,
		p.CreatedAt,
		p.StartedAt,
		p.CompletedAt,
	}
}

type cmd struct {
	cfg Config
	db  database
}

func newCmd(cfg Config, db database) *cmd {
	return &cmd{cfg: cfg, db: db}
}

func (c *cmd) run(n int) (string, error) {
	if err := c.db.connect(); err != nil {
		return "", fmt.Errorf("DB connect: %v", err)
	}

	if err := c.db.init(); err != nil {
		return "", fmt.Errorf("DB init: %v", err)
	}

	// Generate packages
	pkgs := make([]pkg, n)
	for i := range n {
		pkgs[i] = gen(i + 1)
	}

	count, err := c.db.batchInsertPkgs(pkgs)
	if err != nil {
		return "", fmt.Errorf("insert packages: %v", err)
	}

	return fmt.Sprintf("%d packages inserted!\n", count), nil
}

type mysqlDB struct {
	cfg  mysql.Config
	conn *sql.DB
}

func (d *mysqlDB) connect() error {
	// Get a database handle.
	conn, err := sql.Open("mysql", d.cfg.FormatDSN())
	if err != nil {
		return err
	}
	d.conn = conn

	return nil
}

func (d *mysqlDB) init() error {
	if _, err := d.conn.Exec("DELETE FROM package;"); err != nil {
		return fmt.Errorf("delete packages: %v", err)
	}

	return nil
}

func (d *mysqlDB) batchInsertPkgs(pkgs []pkg) (int, error) {
	var count int
	batch := make([]pkg, 0, batchSize)

	for _, p := range pkgs {
		batch = append(batch, p)
		count++

		if count%batchSize == 0 {
			if _, err := d.insertPkgs(batch); err != nil {
				return 0, fmt.Errorf("batch: %v", err)
			}
			batch = make([]pkg, 0, batchSize)
		}
	}

	if len(batch) > 0 {
		if _, err := d.insertPkgs(batch); err != nil {
			return 0, fmt.Errorf("batch: %v", err)
		}
	}

	return count, nil
}

func (d *mysqlDB) insertPkgs(pkgs []pkg) (int, error) {
	args := make([]any, 0, 10*len(pkgs))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	q := "INSERT INTO package VALUES "
	t := "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"

	for _, p := range pkgs {
		if p.ID == "" {
			break
		}

		q += t
		args = append(args, p.toSlice()...)
	}

	// Trim final comma.
	q = q[0 : len(q)-1]
	q += ";"

	r, err := d.conn.ExecContext(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("insert packages: %v", err)
	}

	count, err := r.RowsAffected()
	if err != nil {
		count = int64(len(pkgs))
	}

	return int(count), nil
}

func gen(i int) pkg {
	const doneStatus string = "2"
	return pkg{
		ID:          strconv.Itoa(i),
		Name:        fmt.Sprintf("DPJ-SIP-%s.tar", id()),
		WorkflowID:  fmt.Sprintf("processing-workflow-%s", id()),
		RunID:       uuid.New(),
		AIPID:       uuid.New(),
		LocationID:  uuid.New(),
		Status:      doneStatus,
		CreatedAt:   time.Date(2019, 11, 21, 17, 36, 10, 0, time.UTC),
		StartedAt:   time.Date(2019, 11, 21, 17, 36, 11, 0, time.UTC),
		CompletedAt: time.Date(2019, 11, 21, 17, 42, 12, 0, time.UTC),
	}
}

func id() string {
	return uuid.New().String()
}
