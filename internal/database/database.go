package database

import (
	"auth/api/route"
	"auth/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var (
	database   string
	password   string
	username   string
	port       string
	host       string
	schema     string
	dbInstance *sql.DB
)

func connectDB() *sql.DB {
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

func InitServer() *http.Server {
	loadEnviroment()

	portToInt, _ := strconv.Atoi(os.Getenv("PORT"))

	server := &route.Server{
		Server: model.Server{
			Port: portToInt,
			Db:   connectDB(),
		},
	}

	serverConfig := &http.Server{
		Addr:         fmt.Sprintf(":%d", server.Port),
		Handler:      server.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return serverConfig
}

func loadEnviroment() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading env file")
	}

	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port = os.Getenv("DB_PORT")
	host = os.Getenv("DB_HOST")
	schema = os.Getenv("DB_SCHEMA")
}
