package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"prova-2024/config"
	"prova-2024/handler"
	"prova-2024/helper"
	"prova-2024/model"
	"prova-2024/repository"
	"prova-2024/router"
	"prova-2024/service"
	"time"
)

func main() {
	db := config.DatabaseConnection()
	validate := validator.New()

	err := db.Table("vehicle").AutoMigrate(&model.Vehicle{})
	if err != nil {
		return
	}

	//Init Repository
	vehicleRepository := repository.NewVehicleRepositoryImpl(db)
	accessoryRepository := repository.NewAccessoryRepositoryImpl(db)

	//Init Service
	vehicleService := service.NewVehicleServiceImpl(vehicleRepository, validate)
	accessoryService := service.NewAccessoryServiceImpl(accessoryRepository, validate)

	//Init handler
	vehicleHandler := handler.NewVehicleHandler(vehicleService)
	accessoryHandler := handler.NewAccessoryHandler(accessoryService)

	//Router
	routes := router.NewRouter(vehicleHandler, accessoryHandler)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
