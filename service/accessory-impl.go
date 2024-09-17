package service

import (
	"github.com/go-playground/validator/v10"
	"prova-2024/data/request"
	"prova-2024/helper"
	"prova-2024/model"
	"prova-2024/repository"
)

type AccessoryServiceImpl struct {
	AccessoryRepository repository.AccessoryRepository
	Validate            *validator.Validate
}

func NewAccessoryServiceImpl(accessoryRepository repository.AccessoryRepository, validate *validator.Validate) AccessoryService {
	return &AccessoryServiceImpl{
		AccessoryRepository: accessoryRepository,
		Validate:            validate,
	}
}

func (s AccessoryServiceImpl) Create(accessory request.CreateAccessoryRequest) {
	err := s.Validate.Struct(accessory)
	helper.ErrorPanic(err)
	accessoryModel := model.Accessory{
		VehicleId: accessory.VehicleId,
		Name:      accessory.Name,
	}
	err = s.AccessoryRepository.Save(accessoryModel)
	if err != nil {
		return
	}
}

func (s AccessoryServiceImpl) Delete(accessory request.DeleteAccessoryRequest) {
	err := s.Validate.Struct(accessory)
	accessoryModel := model.Accessory{
		Id:        accessory.AccessoryId,
		VehicleId: accessory.VehicleId,
	}

	err = s.AccessoryRepository.Delete(accessoryModel)
	if err != nil {
		return
	}
}
