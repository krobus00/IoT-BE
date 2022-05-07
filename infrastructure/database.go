package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	SqlxDB *sqlx.DB
	DB     *sql.DB
}

func NewDatabase(env Env) Database {
	USER := env.DBUsername
	PASS := env.DBPassword
	HOST := env.DBHost
	PORT := env.DBPort
	DBNAME := env.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Connection setup failed: %v", err)
	}

	return Database{
		DB:     conn.DB,
		SqlxDB: conn,
	}
}
