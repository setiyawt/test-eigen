package main

import (
	"log"
	"myproject/api"
	"myproject/db"
	"myproject/model"
	repo "myproject/repository"
	"myproject/service"
)

func main() {
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "eigen",
		Port:         5432,
	}

	dbConn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	err = db.SQLExecute(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	userRepo := repo.NewUserRepo(dbConn)
	sessionRepo := repo.NewSessionRepo(dbConn)
	bookRepo := repo.NewBookRepo(dbConn)
	borrowRepo := repo.NewBorrowRepo(dbConn)
	penaltiesRepo := repo.NewPenaltiesRepo(dbConn)

	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	bookService := service.NewBookService(bookRepo)
	borrowService := service.NewBorrowService(borrowRepo)
	penaltiesService := service.NewPenaltiesService(penaltiesRepo)

	mainAPI := api.NewAPI(userService, sessionService, bookService, borrowService, penaltiesService)
	mainAPI.Start()
}
