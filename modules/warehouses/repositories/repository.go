package repositories

import "github.com/ahmadkarlam/go-shipping/modules/warehouses/models"

type WarehouseRepository interface {
	GetAll() ([]models.Warehouse, error)
	DecreaseStock(warehouse models.Warehouse) error
}
