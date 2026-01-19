package GolangDBMySQL

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}