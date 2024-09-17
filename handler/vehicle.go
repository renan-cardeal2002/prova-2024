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

type VehicleHandler struct {
	vehicleService service.VehicleService
}

func NewVehicleHandler(service service.VehicleService) *VehicleHandler {
	return &VehicleHandler{vehicleService: service}
}

func (h *VehicleHandler) Create(ctx *gin.Context) {
	createVehicleRequest := request.CreateVehicleRequest{}
	err := ctx.ShouldBindJSON(&createVehicleRequest)
	helper.ErrorPanic(err)

	h.vehicleService.Create(createVehicleRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (h *VehicleHandler) Update(ctx *gin.Context) {
	updateVehicleRequest := request.UpdateVehicleRequest{}
	err := ctx.ShouldBindJSON(&updateVehicleRequest)
	helper.ErrorPanic(err)

	vehicleId := ctx.Param("vehicleId")
	id, err := strconv.Atoi(vehicleId)
	helper.ErrorPanic(err)

	updateVehicleRequest.Id = id

	h.vehicleService.Update(updateVehicleRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (h *VehicleHandler) Delete(ctx *gin.Context) {
	vehicleId := ctx.Param("vehicleId")
	id, err := strconv.Atoi(vehicleId)
	helper.ErrorPanic(err)
	h.vehicleService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (h *VehicleHandler) FindById(ctx *gin.Context) {
	vehicleId := ctx.Param("vehicleId")
	id, err := strconv.Atoi(vehicleId)
	helper.ErrorPanic(err)

	vehicleResponse := h.vehicleService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   vehicleResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *VehicleHandler) FindAll(ctx *gin.Context) {
	vehicleResponse := h.vehicleService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   vehicleResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
