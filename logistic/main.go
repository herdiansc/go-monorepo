package main

import (
	"errors"
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
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Verifying Token.....")

		authVerifyURL := fmt.Sprintf("http://%s:%s/verify", os.Getenv("AUTH_HOST"), os.Getenv("AUTH_PORT"))
		hr, err := http.NewRequest("POST", authVerifyURL, nil)
		if err != nil {
			log.Printf("Failed to create new http request: %+v\n", err.Error())
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		hr.Header.Add("Content-Type", "application/json")
		hr.Header.Add("Authorization", r.Header.Get("Authorization"))

		client := &http.Client{}
		resp, err := client.Do(hr)
		if err != nil {
			log.Printf("Failed to perform http request: %+v\n", err.Error())
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if resp.StatusCode != http.StatusOK {
			log.Printf("aa Failed to verify token\n")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// @title Logistic Service
// @version 1.0
// @description This is a logistic service
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
	err = DB.AutoMigrate(&models.Logistic{})
	if err == nil && DB.Migrator().HasTable(&models.Logistic{}) {
		if err := DB.First(&models.Logistic{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			firstData := models.Logistic{
				LogisticName:    "JNE",
				Amount:          10000,
				DestinationName: "JAKARTA",
				OriginName:      "BANDUNG",
				Duration:        "2-4",
			}
			DB.Create(&firstData)
		}
	}

	port := os.Getenv("SERVICE_PORT")
	httpServer := http.NewServeMux()
	httpServer.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)), //The url pointing to API definition
	))
	logisticHandlers := handlers.NewLogisticHandler(DB)
	httpServer.Handle("GET /logistics", Authenticate(http.HandlerFunc(logisticHandlers.List)))
	httpServer.Handle("GET /logistics/{uuid}", Authenticate(http.HandlerFunc(logisticHandlers.GetLogisticByUUID)))

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), httpServer)
	if err != nil {
		panic(err)
	}
}
