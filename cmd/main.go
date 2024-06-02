package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	"ahbcc/cmd/migrations"
	"ahbcc/internal/setup"
)

const migrationsDir string = "./migrations"

func main() {
	ctx := context.Background()
	conn := setup.Init(pgx.Connect(ctx, databaseURL()))
	defer conn.Close(ctx)

	runMigrations := migrations.MakeRun(conn)
	err := runMigrations(ctx, migrationsDir)
	if err != nil {
		log.Fatal(err)
	}
}

func databaseURL() string {
	dbUser := os.Getenv("POSTGRES_DB_USER")
	dbPass := os.Getenv("POSTGRES_DB_PASS")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	dbHost := os.Getenv("POSTGRES_DB_HOST")
	dbPort := os.Getenv("POSTGRES_DB_PORT")

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbHost, dbPort)
}
