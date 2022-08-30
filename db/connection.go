package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sebarray/wiselink/config"
)

func ConnectionDB() *sql.DB {
	config.Loadenv()
	var err error
	driver := "mysql"
	cadena := os.Getenv("STRING_CONNECTION")

	conn, err := sql.Open(driver, cadena)
	if err != nil {
		fmt.Println(err.Error())
	}

	return conn

}

func Checkconnection() bool {
	err := ConnectionDB().Ping()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
