package data

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

// OpenDB is response for opening a connection to an underlying SQL DB Pool.
func OpenDB(infoLog *log.Logger) (*sql.DB, error) {
	// In production (or just Dev) this would only exist in our env file. Set for ease of reviewer.
	err := os.Setenv("POSTGRES_PASSWORD_URLSHORTENER", "LemonHotspotAppleRider2134")
	if err != nil {
		return nil, err
	}
	password := os.Getenv("POSTGRES_PASSWORD_URLSHORTENER")
	dsn := fmt.Sprintf("postgres://postgres:%s@localhost/postgres?sslmode=disable", password)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetConnMaxIdleTime(15 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	infoLog.Println("successfully connected to db pool")
	return db, nil
}
