package repository

import (
	"errors"
	"gorm.io/gorm"
	"prova-2024/data/request"
	"prova-2024/helper"
	"prova-2024/model"
)

type VehicleRepositoryImpl struct {
	Db *gorm.DB
}

func NewVehicleRepositoryImpl(Db *gorm.DB) VehicleRepository {
	return &VehicleRepositoryImpl{Db: Db}
}

func (r VehicleRepositoryImpl) Save(vehicle model.Vehicle) {
	result := r.Db.Create(&vehicle)
	helper.ErrorPanic(result.Error)
}

func (r VehicleRepositoryImpl) Update(vehicle model.Vehicle) {
	var updateVehicle = request.UpdateVehicleRequest{
		Id:    vehicle.Id,
		Year:  vehicle.Year,
		Model: vehicle.Model,
		Plate: vehicle.Plate,
	}
	result := r.Db.Model(&vehicle).Updates(updateVehicle)
	helper.ErrorPanic(result.Error)
}

func (r VehicleRepositoryImpl) Delete(vehicleId int) {
	var vehicle model.Vehicle
	result := r.Db.Where("id = ?", vehicleId).Delete(&vehicle)
	helper.ErrorPanic(result.Error)
}

func (r VehicleRepositoryImpl) FindById(vehicleId int) (model.Vehicle, error) {
	var vehicle model.Vehicle
	result := r.Db.Find(&vehicle, vehicleId)
	if result != nil {
		return vehicle, nil
	} else {
		return vehicle, errors.New("tag is not found")
	}
}

func (r VehicleRepositoryImpl) FindAll() []model.Vehicle {
	var vehicle []model.Vehicle
	results := r.Db.Find(&vehicle)
	helper.ErrorPanic(results.Error)
	return vehicle
}
