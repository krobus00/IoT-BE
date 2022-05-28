package infrastructure

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
)

func NewDatabase(env Env) kro_pkg.Database {
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
	wrapper := kro_pkg.NewDatabase(conn)
	return *wrapper
}
