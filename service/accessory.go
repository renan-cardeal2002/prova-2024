package service

import (
	"prova-2024/data/request"
)

type AccessoryService interface {
	Create(accessory request.CreateAccessoryRequest)
	Delete(accessory request.DeleteAccessoryRequest)
}
