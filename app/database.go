package app

import (
	"database/sql"
	"time"

	"github.com/Benyamin206/backend-clean-arch-go/helper"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func DBConnection() *sql.DB {

	db, err := sql.Open("pgx", "postgres://postgres:example@localhost:5432/example?sslmode=disable")

	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)

	return db
}
