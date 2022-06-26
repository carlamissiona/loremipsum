package database

import (
	"database/sql"
	_ "encoding/json"
	_ "fmt"
	"log"
	_ "net/http"
	"os"
	_ "os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupDatabase() ( *sql.DB ) {
 
	err := godotenv.Load(".environ_development")
	if err != nil {
		log.Fatalf("failed reading env file: %v", err)
	}
	log.Println("failed reading env file: %v")
	db_info := os.Getenv("POSTGRES_URL")
	Database, err := sql.Open("postgres", db_info)
	if err != nil {
		panic(err)
	}

	defer Database.Close()

	err = Database.Ping()
	if err != nil {
		log.Fatalf("failed No DB connection %v", err)
	}
  log.Println("db_instance graph")
  return  Database
}
