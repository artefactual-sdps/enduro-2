package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const (
	// defaultSize is the default number of records to generate.
	defaultSize = 1000

	// batchSize is the number of rows to insert with a single insert query.
	batchSize = 100
)

type database interface {
	connect() error
	init() error
	batchInsertPkgs([]pkg) (int, error)
	insertPkgs([]pkg) (int, error)
	getPackages(ctx context.Context, limit, cursor int) ([]pkg, error)
	insertActions(context.Context, []pkg) ([]int, error)
}

type Config struct{}

func main() {
	args := os.Args

	n := defaultSize
	if len(args) > 1 {
		i, err := strconv.Atoi(args[1])
		if err != nil {
			usage()
			fmt.Println("Error: first argument must be a number")
			os.Exit(1)
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
Generates random package data for testing.

USAGE: go run main.go [COUNT]

COUNT - Number of packages to be created (default: %d)

`,
		defaultSize,
	)
}

type pkg struct {
	ID          int
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
	var cursor int

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Add preservation actions and tasks.
	for {
		pkgs, err := c.db.getPackages(ctx, batchSize, cursor)
		if err != nil {
			return "", err
		}

		_, err = c.db.insertActions(ctx, pkgs)
		if err != nil {
			return "", err
		}

		if len(pkgs) < batchSize {
			break
		}

		lastPkg := pkgs[len(pkgs)-1]
		cursor = lastPkg.ID
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
		if p.ID == 0 {
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

func (d *mysqlDB) getPackages(ctx context.Context, limit, cursor int) ([]pkg, error) {
	var res []pkg
	args := []any{}

	q := "SELECT id FROM package"
	if cursor > 0 {
		q += " WHERE id > ?"
		args = append(args, cursor)
	}

	q += " LIMIT ?"
	args = append(args, limit)

	r, err := d.conn.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("get packages: %v", err)
	}
	defer r.Close()

	for r.Next() {
		var p pkg
		if err := r.Scan(&p.ID); err != nil {
			return nil, fmt.Errorf("set package ID: %v", err)
		}
		res = append(res, p)
	}

	return res, nil
}

func (d *mysqlDB) insertActions(ctx context.Context, pkgs []pkg) ([]int, error) {
	var args []any

	pkgIDs := make([]int, len(pkgs))

	if len(pkgs) == 0 {
		return nil, nil
	}

	q := `
INSERT INTO preservation_action (
	workflow_id,
	type,
	status,
	started_at,
	completed_at,
	package_id
) VALUES`

	for i, p := range pkgs {
		args = append(args,
			fmt.Sprintf("processing-workflow-%s", id()), // workflow_id.
			1, // type: "Create AIP".
			2, // status: "Done".
			time.Date(2019, 11, 21, 17, 36, 10, 0, time.UTC), // started_at.
			time.Date(2019, 11, 21, 17, 37, 8, 0, time.UTC),  // completed_at.
			p.ID,
		)
		q += " (?, ?, ?, ?, ?, ?),"
		pkgIDs[i] = p.ID
	}

	// Replace final comma with a semicolon.
	q = strings.TrimSuffix(q, ",") + ";"

	_, err := d.conn.ExecContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("insert preservation actions: %v", err)
	}

	// Return inserted actions.
	args = []any{}
	q = "SELECT FROM preservation_action WHERE package_id IN ("
	for _, id := range pkgIDs {
		q += "?,"
		args = append(args, id)
	}
	q += q[0:len(q)-1] + ");"

	r, err := d.conn.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("select preservation actions: %v", err)
	}

	var ids []int
	for r.Next() {
		var id int
		r.Scan(&id)
		ids = append(ids, id)
	}

	return ids, nil
}

func gen(i int) pkg {
	const doneStatus string = "2"
	return pkg{
		ID:          i,
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
