package helpers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const MYSQL_DB = "mysql"

func BuildMysqlConnUrl() string {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", user, pass, host, port, dbName)
}

func StartMysqlDb() *sql.DB {
	db, err := sql.Open(MYSQL_DB, BuildMysqlConnUrl())
	if err != nil {
		panic(err)
	}

	return db
}

