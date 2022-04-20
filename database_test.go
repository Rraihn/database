package go_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenconnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang?parseTime=true")
	if err != nil {
		panic(err)
		defer db.Close()
	}
}
