package repository

import "prova-2024/model"

type VehicleRepository interface {
	Save(vehicle model.Vehicle)
	Update(vehicle model.Vehicle)
	Delete(vehicleId int)
	FindById(vehicleId int) (tags model.Vehicle, err error)
	FindAll() []model.Vehicle
}
