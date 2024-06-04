package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"ahbcc/cmd/migrations"
	"ahbcc/internal/setup"

	"github.com/jackc/pgx/v5/pgxpool"
)

const migrationsDir string = "./migrations"

func main() {
	ctx := context.Background()
	pool := setup.Init(pgxpool.New(ctx, databaseURL()))
	defer pool.Close()

	runMigrations := migrations.MakeRun(pool)
	err := runMigrations(ctx, migrationsDir)
	if err != nil {
		log.Fatal(err)
	}
}

func databaseURL() string {
	dbUser := os.Getenv("POSTGRES_DB_USER")
	dbPass := os.Getenv("POSTGRES_DB_PASS")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	dbPort := os.Getenv("POSTGRES_DB_PORT")

	return fmt.Sprintf("user=%s password=%s dbname=%s host=postgres_db port=%s sslmode=disable", dbUser, dbPass, dbName, dbPort)
}
