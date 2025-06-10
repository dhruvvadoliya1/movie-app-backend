package database

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql" // import mysql if it is used
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Seeder struct {
	Db *goqu.Database
}

func NewSeeder(db *goqu.Database) *Seeder {
	return &Seeder{Db: db}
}
