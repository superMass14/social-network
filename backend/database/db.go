package database

import (
	q "backend/database/query"
	"backend/utils"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
}

var DB *Database

func init() {
	db, err := sql.Open("sqlite3", "./database/social_network.db")
	if err != nil {
		log.Println("Error opening database:", err)
		os.Exit(1)
	}
	err = MigrateDB(db)
	if err != nil {
		log.Println("Error migrating database:", err)
		os.Exit(1)
	}
	//seed.CreateTable(db)
	//seed.InsertData(db)
	log.Println("Database opened")
	DB = &Database{db}
}

func NewDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "../backend/database/social_network.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	return db
}

func MigrateDB(db *sql.DB) error {
	utils.ClearScreen()
	log.Println("Migrating database...")
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %v", err)
	}

	// Get the absolute path to the migrations directory
	absPath, err := filepath.Abs("../backend/database/migrations/sqlite")
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Run migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath,
		"sqlite3", driver)

	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("an error occurred while syncing the database: %v", err)
	}

	log.Println("Migration completed successfully")
	return nil
}

func (d *Database) Insert(table string, data any) error {
	query, err := q.InsertQuery(table, data)
	if err != nil {
		return fmt.Errorf("error creating insert query: %v", err)
	}
	prep, err := d.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing insert query: %v", err)
	}
	_, err = prep.Exec()

	return err
}

func (d *Database) Delete(table string, where q.WhereOption) error {

	query := q.DeleteQuery(table, where)
	stmt, err := d.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing delete query: %v", err)
	}
	_, err = stmt.Exec()
	return err
}

func (d *Database) Update(table string, object any, where q.WhereOption) error {
	var err error
	query, err := q.UpdateQuery(table, object, where)
	if err != nil {
		return fmt.Errorf("error creating update query: %v", err)
	}
	stmt, err := d.Prepare(query)

	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

func (d *Database) GetOneFrom(table string, where q.WhereOption) (*sql.Row, error) {

	query := q.SelectOneFrom(table, where)
	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow()
	return row, nil
}

func (d *Database) GetAllFrom(table string, where q.WhereOption, orderby string, limit []int) (*sql.Rows, error) {
	var query string
	if where == nil {
		query = q.SelectAllFrom(table, orderby, limit)
	} else {
		query = q.SelectAllWhere(table, where, orderby, limit)
	}
	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing select query: %v", err)
	}
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Database) GetAllAndJoin(table string, j []q.JoinCondition, where q.WhereOption, orderby string, limit []int) (*sql.Rows, error) {
	var query = q.SelectWithJoinQuery(table, j, where, orderby, limit)

	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Database) GetCount(table string, where q.WhereOption) (*sql.Row, error) {
	var query = q.GetCountQuery(table, where)

	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow()

	return row, nil
}
