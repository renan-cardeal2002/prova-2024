package service

import (
	"github.com/go-playground/validator/v10"
	"prova-2024/data/request"
	"prova-2024/data/response"
	"prova-2024/helper"
	"prova-2024/model"
	"prova-2024/repository"
)

type VehicleServiceImpl struct {
	VehicleRepository repository.VehicleRepository
	Validate          *validator.Validate
}

func NewVehicleServiceImpl(vehicleRepository repository.VehicleRepository, validate *validator.Validate) VehicleService {
	return &VehicleServiceImpl{
		VehicleRepository: vehicleRepository,
		Validate:          validate,
	}
}

func (s VehicleServiceImpl) Create(vehicle request.CreateVehicleRequest) {
	err := s.Validate.Struct(vehicle)
	helper.ErrorPanic(err)
	vehicleModel := model.Vehicle{
		Year:  vehicle.Year,
		Model: vehicle.Model,
		Plate: vehicle.Plate,
	}
	s.VehicleRepository.Save(vehicleModel)
}

func (s VehicleServiceImpl) Update(vehicle request.UpdateVehicleRequest) {
	vehicleData, err := s.VehicleRepository.FindById(vehicle.Id)
	helper.ErrorPanic(err)
	vehicleData.Year = vehicle.Year
	vehicleData.Model = vehicle.Model
	vehicleData.Plate = vehicle.Plate
	s.VehicleRepository.Update(vehicleData)
}

func (s VehicleServiceImpl) Delete(vehicleId int) {
	s.VehicleRepository.Delete(vehicleId)
}

func (s VehicleServiceImpl) FindById(vehicleId int) response.VehicleResponse {
	vehicleData, err := s.VehicleRepository.FindById(vehicleId)
	helper.ErrorPanic(err)

	vehicleResponse := response.VehicleResponse{
		Id:    vehicleData.Id,
		Year:  vehicleData.Year,
		Model: vehicleData.Model,
		Plate: vehicleData.Plate,
	}
	return vehicleResponse
}

func (s VehicleServiceImpl) FindAll() []response.VehicleResponse {
	result := s.VehicleRepository.FindAll()

	var vehicles []response.VehicleResponse
	for _, value := range result {
		vehicle := response.VehicleResponse{
			Id:    value.Id,
			Year:  value.Year,
			Model: value.Model,
			Plate: value.Plate,
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles
}
