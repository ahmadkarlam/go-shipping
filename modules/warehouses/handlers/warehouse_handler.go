package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ahmadkarlam/go-shipping/common/resolver"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/dto"
)

type WarehouseHandler struct {
	resolver resolver.Resolver
}

func NewWarehouseHandler(resolver resolver.Resolver) WarehouseHandler {
	return WarehouseHandler{resolver: resolver}
}

// GetAll godoc
// @Summary get all warehouse
// @Description get all data warehouse
// @Tag Warehouse
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {string} string "Internal Server Error"
// @Router /warehouse [get]
func (h *WarehouseHandler) GetAll(ctx *gin.Context) {
	warehouses, err := h.resolver.WarehouseService.GetAllWarehouse()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": warehouses,
	})
}

// SendVaccineToLocation godoc
// @Summary send vaccine
// @Description send vaccine to given location
// @Tag Warehouse
// @Accept  json
// @Produce  json
// @Param request body dto.SendVaccineToLocationRequest true "Request Body"
// @Success 200 {string} string "OK"
// @Failure 500 {string} string "Internal Server Error"
// @Router /warehouse/send-vaccine [post]
func (h *WarehouseHandler) SendVaccineToLocation(ctx *gin.Context) {
	var request dto.SendVaccineToLocationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sendingCost, err := h.resolver.WarehouseService.FindNearbyWarehouse(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "sending vaccine",
		"data": gin.H{
			"from_warehouse": sendingCost.From,
			"distance":       sendingCost.Distance,
			"cost":           sendingCost.Cost,
			"time_took": gin.H{
				"days":  sendingCost.Day,
				"hours": sendingCost.Hour,
			},
		},
	})
}
