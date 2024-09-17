package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prova-2024/data/request"
	"prova-2024/data/response"
	"prova-2024/helper"
	"prova-2024/service"
	"strconv"
)

type AccessoryHandler struct {
	accessoryService service.AccessoryService
}

func NewAccessoryHandler(service service.AccessoryService) *AccessoryHandler {
	return &AccessoryHandler{accessoryService: service}
}

func (h *AccessoryHandler) Create(ctx *gin.Context) {
	createAccessoryRequest := request.CreateAccessoryRequest{}
	err := ctx.ShouldBindJSON(&createAccessoryRequest)
	helper.ErrorPanic(err)

	h.accessoryService.Create(createAccessoryRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (h *AccessoryHandler) Delete(ctx *gin.Context) {
	vehicleId := ctx.Param("vehicleId")
	idVehicle, err := strconv.Atoi(vehicleId)
	helper.ErrorPanic(err)

	accessoryId := ctx.Param("accessoryId")
	idAccessory, err := strconv.Atoi(accessoryId)
	helper.ErrorPanic(err)

	accessoryRequest := request.DeleteAccessoryRequest{
		VehicleId:   idVehicle,
		AccessoryId: idAccessory,
	}

	h.accessoryService.Delete(accessoryRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
