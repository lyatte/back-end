package main

import (
	"back_end/graph"
	"back_end/graph/generated"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "5000"

func main() {
	//db := pg.Connect(&pg.Options{
	//	Addr:     ":5432",
	//	User:     "postgres",
	//	Password: "kimkim12",
	//	Database: "YouRJube",
	//})

	opt, err := pg.ParseURL("postgres://cqopwwkufacewc:cb14826403fb985606175b5376e9a74d5082277d9c34dcbe290c127eb9b0fd69@ec2-50-16-198-4.compute-1.amazonaws.com:5432/ddipcc3o4af2rn")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
