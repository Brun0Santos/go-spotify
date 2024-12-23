package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Brun0Santos/api-spotify/internal/handler"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}

	dbCon := os.Getenv("DB_CON")
	appPort := os.Getenv("APP_PORT")

	db, err := sql.Open("postgres", dbCon)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}

	fmt.Println("Server is running....")

	handlerArtists := handler.NewArtistsGet()

	mux := http.NewServeMux()
	mux.Handle("GET /api/v1/artists/{uid}", handlerArtists)

	if err := http.ListenAndServe(appPort, mux); err != nil {
		log.Fatal(err)
	}
}
