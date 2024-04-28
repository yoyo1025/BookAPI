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
	bookRepository := repository.NewBookRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	userController := controller.NewUserController(userUsecase)
	bookController := controller.NewBookController(bookUsecase)
	e := router.NewRouter(userController, bookController)
	e.Logger.Fatal(e.Start(":8080"))
}