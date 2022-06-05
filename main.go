package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Aditya170700/gorestfulapi/app"
	"github.com/Aditya170700/gorestfulapi/controller"
	"github.com/Aditya170700/gorestfulapi/helper"
	"github.com/Aditya170700/gorestfulapi/middleware"
	"github.com/Aditya170700/gorestfulapi/repository"
	"github.com/Aditya170700/gorestfulapi/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
