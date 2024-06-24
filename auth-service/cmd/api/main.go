package main

import (
	"auth-service/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var tries int64 = 0

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres")
	}
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	log.Println("Starting authentication service on port: ", webPort)
	log.Fatal(srv.ListenAndServe())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dns := os.Getenv("DSN")
	for {
		connection, err := openDB(dns)
		if err != nil {
			log.Println("Postgres not yet ready...")
			tries++
		} else {
			log.Println("Connected to Postgres")
			return connection
		}
		if tries > 10 {
			log.Println(err)
			return nil
		}
		log.Println("Backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
