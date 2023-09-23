package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	db := db.NewDB()

	// User related components
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	// Task related components
	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	// Blog related components
	blogValidator := validator.NewBlogValidator()
	blogRepository := repository.NewBlogRepository(db)
	blogUsecase := usecase.NewBlogUsecase(blogRepository, blogValidator)
	blogController := controller.NewBlogController(blogUsecase)

	// Shop related components
	shopValidator := validator.NewShopValidator()
	shopRepository := repository.NewShopRepository(db)
	shopUsecase := usecase.NewShopUsecase(shopRepository, shopValidator)
	shopController := controller.NewShopController(shopUsecase)

	e := router.NewRouter(userController, taskController, blogController, shopController) // Modify the NewRouter function to accept the shopController
	e.Logger.Fatal(e.Start(":8080"))
}
