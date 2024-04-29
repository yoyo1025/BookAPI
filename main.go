package main

import (
	"BookAPI/controller"
	"BookAPI/db"
	"BookAPI/repository"
	"BookAPI/router"
	"BookAPI/usecase"
	"BookAPI/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	bookValidator := validator.NewBookValidator()
	userRepository := repository.NewUserRepository(db)
	bookRepository := repository.NewBookRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	bookUsecase := usecase.NewBookUsecase(bookRepository, bookValidator)
	userController := controller.NewUserController(userUsecase)
	bookController := controller.NewBookController(bookUsecase)
	e := router.NewRouter(userController, bookController)
	e.Logger.Fatal(e.Start(":8080"))
}