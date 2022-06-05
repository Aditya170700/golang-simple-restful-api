package app

import (
	"github.com/Aditya170700/gorestfulapi/controller"
	"github.com/Aditya170700/gorestfulapi/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:category", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:category", categoryController.Update)
	router.DELETE("/api/categories/:category", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
