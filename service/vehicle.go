package service

import (
	"prova-2024/data/request"
	"prova-2024/data/response"
)

type VehicleService interface {
	Create(vehicle request.CreateVehicleRequest)
	Update(vehicle request.UpdateVehicleRequest)
	Delete(vehicleId int)
	FindById(vehicleId int) response.VehicleResponse
	FindAll() []response.VehicleResponse
}
