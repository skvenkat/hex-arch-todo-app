package main

import (
	"database/sql"
	"embed"
	"fmt"

	v4 "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/skvenkat/hex-arch-todo-app/helpers"
)

var (
	//go:embed sql/*.sql
	migFs embed.FS
)

func main() {
	fmt.Println("Migration process initated...")
	db, err := sql.Open("mysql", helpers.BuildMysqlConnUrl())
	if err != nil {
		panic(err)
	}

	//driver, err := mysql.NewFromDB(db)
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	dirs, err := migFs.ReadDir("sql")
	if err != nil {
		panic(fmt.Sprintf("error reading migration files from embed :%v", err))
	} else {
		fmt.Println("List of Migrations found :")
		for _, d := range dirs {
			fmt.Sprintf(" - %s\n", d.Name())
		}
		fmt.Println("End of List")
	}

	/*
	embedSource := &migration.EmbedMigrationSource{
		EmbedFS: migFs,
		Dir: "sql", 
	}
	*/

	// Run all up migrations
	//applied, err := migration.Migrate(driver, embedSource, migration.Up, 0)
	
	m, err := v4.NewWithDatabaseInstance(
		"file:///sql/*.sql", "mysql", driver,
	)
	m.Up()
	if err != nil {
		panic(fmt.Sprintf("Error applying migrations: %s", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("Last applied %d: ", m))
	}
	fmt.Println("Migration Process Ended...")
}
