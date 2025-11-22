package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/esteveseverson/go-rest-rbac/dbconfig"
	"github.com/esteveseverson/go-rest-rbac/internal/handlers"
	"github.com/esteveseverson/go-rest-rbac/internal/routes"
	"github.com/esteveseverson/go-rest-rbac/internal/store"
	"github.com/esteveseverson/go-rest-rbac/serverconfig"
)

func main() {
	config, err := serverconfig.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	db := dbconfig.ConnectDB(config.DatabaseURL)
	defer db.Close()

	queries := store.New(db)

	handler := handlers.NewHandler(db, queries)

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, handler)

	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	fmt.Printf("Server runing on Port%s\n", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
