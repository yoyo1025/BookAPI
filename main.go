package main

import (
	"BookAPI/controller"
	"BookAPI/db"
	"BookAPI/repository"
	"BookAPI/router"
	"BookAPI/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewuUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}