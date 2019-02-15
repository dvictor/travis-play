package main

import (
	"database/sql"
	"github.com/cenkalti/backoff"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	connStr := "host=db user=postgres password=password dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Connect error", err)
	}

	{
		b := backoff.NewExponentialBackOff()
		b.MaxElapsedTime = 1 * time.Minute
		err := backoff.Retry(func() error {
			log.Println("DB Ping")
			return db.Ping()
		}, b)
		if err != nil {
			log.Fatal("DB retry failed", err)
		}
	}
	{
		rows, err := db.Query("SELECT usename FROM pg_user")
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var name string
			err = rows.Scan(&name)
			log.Println("Name:", name)
		}
	}

	log.Fatal("bla")
}
