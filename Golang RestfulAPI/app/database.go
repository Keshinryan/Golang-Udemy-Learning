package app

import (
	"database/sql"
	"golang_restful_api/helper"
	"time"
)

func NewDB() *sql.DB{
	db,err:=sql.Open("mysql","root@tcp(localhost:3306)/resfulapi")
	helper.PanicErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10*time.Minute)
	db.SetConnMaxLifetime(60*time.Minute)

	return db
}