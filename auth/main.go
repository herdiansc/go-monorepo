package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/herdiansc/orderfaz/auth/docs"
	"github.com/herdiansc/orderfaz/auth/handlers"
	"github.com/herdiansc/orderfaz/auth/models"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func dbOpen() error {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

// @title Auth Service
// @version 1.0
// @description This is a auth service
// @contact.email herdiansc@gmail.com
// @license.name MIT
// @BasePath /
// @query.collection.format multi
func main() {
	fmt.Println("Running server")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	if err := dbOpen(); err != nil {
		log.Fatalf("Error connecting to the database: %+v\n", err)
		return
	}
	log.Printf("DB connection established: %+v\n", DB)
	DB.AutoMigrate(&models.Auth{})

	port := os.Getenv("SERVICE_PORT")
	httpServer := http.NewServeMux()
	httpServer.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)), //The url pointing to API definition
	))
	authHandlers := handlers.NewAuthHandler(DB)
	httpServer.HandleFunc("POST /register", authHandlers.Register)
	httpServer.HandleFunc("POST /login", authHandlers.Login)
	httpServer.HandleFunc("POST /verify", authHandlers.Verify)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), httpServer)
	if err != nil {
		panic(err)
	}
}
