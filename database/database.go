package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DB  *sql.DB
	err error
	once sync.Once
)

func DBConnection() *sql.DB {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load environment file")
		panic(err)
	} else {
		fmt.Println("Success to load environment file")
	}

	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME"))
	
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Connection to database failed")
		panic(err)
	} else {
		fmt.Println("Connection to database success")
	}
	return DB
}

func DBStart() *sql.DB {
	once.Do(func() {
		DB = DBConnection()
	})
	return DB
}

func DBMigrate() {
	dbParam := DBStart()
	
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	
	if errs != nil {
		fmt.Println("Error while migrating database")
		panic(errs)
	} else {
		fmt.Println("Success migrating database")
	}

	fmt.Println("Applied", n, "migrations!")
}