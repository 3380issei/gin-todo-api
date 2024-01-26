package main

import (
	"todo-api/controller"
	"todo-api/db"
	"todo-api/middleware"
	"todo-api/repository"
	"todo-api/router"
	"todo-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	authMiddleware := middleware.NewAuthMiddleware()
	r := router.NewRouter(userController, taskController, authMiddleware)

	r.Run()
}
