package main

import (
	"log"
	"net/http"

	"myproject/api"
	"myproject/db"
	_ "myproject/docs" // Import generated docs
	"myproject/model"
	repo "myproject/repository"
	"myproject/service"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Library Management API
// @version 1.0
// @description This is a sample API for managing a library.
// @host localhost:1323
// @BasePath /
// @schemes http
func main() {
	// Database credentials
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "eigen",
		Port:         5432,
	}

	// Connect to the database
	dbConn, err := db.Connect(&dbCredential)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Execute SQL scripts
	err = db.SQLExecute(dbConn)
	if err != nil {
		log.Fatalf("Failed to execute SQL scripts: %v", err)
	}

	defer dbConn.Close()

	// Initialize repositories
	userRepo := repo.NewUserRepo(dbConn)
	sessionRepo := repo.NewSessionRepo(dbConn)
	bookRepo := repo.NewBookRepo(dbConn)
	borrowRepo := repo.NewBorrowRepo(dbConn)
	penaltiesRepo := repo.NewPenaltiesRepo(dbConn)

	// Initialize services
	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	bookService := service.NewBookService(bookRepo)
	borrowService := service.NewBorrowService(borrowRepo)
	penaltiesService := service.NewPenaltiesService(penaltiesRepo)

	// Create new API
	mainAPI := api.NewAPI(userService, sessionService, bookService, borrowService, penaltiesService)

	r := chi.NewRouter()

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"),
	))

	// Register API routes
	r.Mount("/", mainAPI.Handler())

	// Start the server
	log.Println("Starting server on :1323")
	if err := http.ListenAndServe(":1323", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
