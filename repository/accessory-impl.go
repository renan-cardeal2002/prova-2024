package repository

import (
	"gorm.io/gorm"
	"prova-2024/helper"
	"prova-2024/model"
)

type AccessoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewAccessoryRepositoryImpl(Db *gorm.DB) AccessoryRepository {
	return &AccessoryRepositoryImpl{Db: Db}
}

func (r AccessoryRepositoryImpl) Save(accessory model.Accessory) error {
	result := r.Db.Create(&accessory)
	helper.ErrorPanic(result.Error)
	return result.Error
}

func (r AccessoryRepositoryImpl) Delete(accessory model.Accessory) error {
	var accessoryResult model.Accessory
	result := r.Db.Where("id = $1 and vehicle_id = $2", accessory.Id, accessory.VehicleId).Delete(&accessoryResult)
	helper.ErrorPanic(result.Error)
	return result.Error
}
