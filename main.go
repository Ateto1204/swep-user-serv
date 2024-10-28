package main

import (
	"log"

	"github.com/Ateto1204/swep-user-serv/internal/infrastructure"
	"github.com/Ateto1204/swep-user-serv/internal/repository"
	"github.com/Ateto1204/swep-user-serv/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := infrastructure.NewDatabase()
	if err != nil {
		panic(err)
	}

	repo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(repo)

	go userUseCase.Run()

	router := infrastructure.NewRouter(userUseCase)
	log.Println("Server Start:")
	router.Run(":8080")
}
