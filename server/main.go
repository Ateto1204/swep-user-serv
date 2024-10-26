package main

import (
	"github.com/Ateto1204/swep-user-serv/server/internal/infrastructure"
	"github.com/Ateto1204/swep-user-serv/server/internal/repository"
	"github.com/Ateto1204/swep-user-serv/server/internal/usecase"
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
	router.Run(":8002")
}
