package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ain-py/go-webserver/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	fmt.Println("hlo")
	portString := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is missing")
	}
	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("cant connect to db", err)
	}
	queries := database.New(conn)

	apiCofg := apiConfig{
		DB: queries,
	}

	if portString == "" {
		log.Fatal("Port not found")
	}
	fmt.Println(portString)
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCofg.handlerCreateUser)
	v1Router.Get("/users", apiCofg.middleWareAuth(apiCofg.handlerGetUser))
	v1Router.Post("/feeds", apiCofg.middleWareAuth(apiCofg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCofg.handlerGetFeeds)
	v1Router.Post("/feed_follows", apiCofg.middleWareAuth(apiCofg.handlerCreateFeedFollow))
	router.Mount("/v1", v1Router)

	log.Printf("Server running ")
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
