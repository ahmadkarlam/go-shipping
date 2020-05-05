package resolver

import (
	"github.com/jinzhu/gorm"

	"github.com/ahmadkarlam/go-shipping/infrastructure/database"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/repositories/mysql"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/services"
)

type Resolver struct {
	WarehouseService services.WarehouseService
}

func NewResolver() Resolver {
	db := database.DBInit()
	warehouseService := initWarehouse(db)
	return Resolver{
		WarehouseService: warehouseService,
	}
}

func initWarehouse(db *gorm.DB) services.WarehouseService {
	repository := mysql.NewWarehouseRepository(db)

	service := services.NewWarehouseService(repository)

	return service
}
