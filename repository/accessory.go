package repository

import "prova-2024/model"

type AccessoryRepository interface {
	Save(accessory model.Accessory) error
	Delete(accessory model.Accessory) error
}
