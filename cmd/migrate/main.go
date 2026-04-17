package main

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	migrations "github/SXsid/url-forge/Migrations"
	"github/SXsid/url-forge/internal/db"
)

func main() {
	dsn := os.Getenv("db_dsn")
	if dsn == "" {
		fmt.Println("connection string is not set")
		os.Exit(-1)
	}
	if err := db.RunMigraton(dsn, migrations.MigratonFs); err != nil {
		fmt.Println(err)
	}
}
