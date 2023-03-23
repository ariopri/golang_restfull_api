package main

import (
	"github.com/ariopri/golang_restfull_api/app"
	"github.com/ariopri/golang_restfull_api/controller"
	"github.com/ariopri/golang_restfull_api/helper"
	"github.com/ariopri/golang_restfull_api/middleware"
	"github.com/ariopri/golang_restfull_api/repository"
	"github.com/ariopri/golang_restfull_api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	// ...

	db := app.NewDB()
	validate := validator.New()
	barangRepository := repository.NewBarangRepository()
	barangService := service.NewBarangService(barangRepository, db, validate)
	barangController := controller.NewBarangController(barangService)

	router := app.NewRouter(barangController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
