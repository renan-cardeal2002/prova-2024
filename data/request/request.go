package request

type CreateVehicleRequest struct {
	Year  int    `json:"year"`
	Model string `json:"model"`
	Plate string `json:"plate"`
}

type UpdateVehicleRequest struct {
	Id    int    `json:"id"`
	Year  int    `json:"year"`
	Model string `json:"model"`
	Plate string `json:"plate"`
}

type CreateAccessoryRequest struct {
	VehicleId int    `json:"vehicleId"`
	Name      string `json:"name"`
}

type DeleteAccessoryRequest struct {
	VehicleId   int `json:"vehicleId"`
	AccessoryId int `json:"accessoryId"`
}
