package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

// from now on, other example will using pgxpool connection instead of pgx connect
func main() {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("GOPGX_DATABASE_URL"))
	// TODO: set connection pooling config

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "SELECT 'Hello World Pool'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nThe SQL greeting with conn pool is '%s'\n", greeting)
}
