package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {
	err := godotenv.Load("../../.env")

	var (
		database   = os.Getenv("DB_DATABASE")
		password   = os.Getenv("DB_PASSWORD")
		username   = os.Getenv("DB_USERNAME")
		port       = os.Getenv("PORT")
		host       = os.Getenv("HOST")
		schema     = os.Getenv("SCHEMA")
		dbInstance *sql.DB
	)

	if err != nil {
		log.Fatal("Error loading env file")
	}

	// Reuse connection

	if dbInstance != nil {
		return dbInstance
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)

	db, err := sql.Open("pgx", connStr)

	if err != nil {
		log.Fatal("Error while connecting to db")
	}

	return db
}

func initServer(port string) *http.Server {
  portToInt, _ := strconv.Atoi(port)

  server := &Server{
	port,
	db: ConnConnectDB(),
  }

  serverConfig := &http.Server{
    Addr: fmt.Sprintf(":%d", server.port),
	Handler: ,
	IdleTimeOut: time.Minute,
	ReadTimeOut: 10 * time.Second,
	WriteTimeout: 30 * time.Second
  }
}
