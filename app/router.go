package app

import (
	"github.com/ariopri/golang_restfull_api/controller"
	"github.com/ariopri/golang_restfull_api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(barangController controller.BarangController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/barangs", barangController.FindAll)
	router.GET("/api/barangs/:BarangId", barangController.FindById)
	router.POST("/api/barangs", barangController.Create)
	router.PUT("/api/barangs/:BarangId", barangController.Update)
	router.DELETE("/api/barangs/:BarangId", barangController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
