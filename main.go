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

	// Favorite related components
	favoriteValidator := validator.NewFavoriteValidator()
	favoriteRepository := repository.NewFavoriteRepository(db)
	favoriteUsecase := usecase.NewFavoriteUsecase(favoriteRepository, shopRepository, userRepository, favoriteValidator)
	favoriteController := controller.NewFavoriteController(favoriteUsecase)

	// Reservation related components
	reservationValidator := validator.NewReservationValidator()
	reservationRepository := repository.NewReservationRepository(db)
	reservationUsecase := usecase.NewReservationUsecase(reservationRepository, reservationValidator)
	reservationController := controller.NewReservationController(reservationUsecase)

	// Initialize the router and start the server
	e := router.NewRouter(userController, taskController, blogController, shopController, favoriteController, reservationController) // Modify to include the reservationController
	e.Logger.Fatal(e.Start(":8080"))
}


