package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"sudutkampus/gorestfulapi/app"
	"sudutkampus/gorestfulapi/controller"
	"sudutkampus/gorestfulapi/helper"
	"sudutkampus/gorestfulapi/middleware"
	"sudutkampus/gorestfulapi/repository"
	"sudutkampus/gorestfulapi/service"

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
