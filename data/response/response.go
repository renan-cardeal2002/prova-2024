package response

type VehicleResponse struct {
	Id    int    `json:"id"`
	Year  int    `json:"year"`
	Model string `json:"model"`
	Plate string `json:"plate"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
