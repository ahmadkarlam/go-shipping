package repositories

import "github.com/ahmadkarlam/go-shipping/modules/warehouses/models"

type WarehouseRepository interface {
	GetAll() ([]models.Warehouse, error)
}
