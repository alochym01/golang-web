package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB type
type DB struct {
	SQL *sql.DB
}

// ConnectSQL function
func ConnectSQL(host string, port int64, username string, password string, dbname string) (*DB, error) {
	dbConn := &DB{}
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		username,
		password,
		host,
		port,
		dbname,
	)

	fmt.Println(dbSource)

	sql, err := sql.Open("mysql", dbSource)
	dbConn.SQL = sql

	if err != nil {
		panic(err)
	}

	return dbConn, err
}
