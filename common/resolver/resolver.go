package resolver

import (
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/repositories/mysql"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/services"
)

type Resolver struct {
	WarehouseService services.WarehouseService
}

func NewResolver() Resolver {
	warehouseService := initWarehouse()
	return Resolver{
		WarehouseService: warehouseService,
	}
}

func initWarehouse() services.WarehouseService {
	repository := mysql.NewWarehouseRepository()

	service := services.NewWarehouseService(repository)

	return service
}
