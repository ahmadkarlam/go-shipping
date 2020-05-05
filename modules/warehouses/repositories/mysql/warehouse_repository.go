package mysql

import (
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/models"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/repositories"
)

var WAREHOUSES = []models.Warehouse{
	{
		Code:  "A",
		Stock: 10,
		X:     5,
		Y:     18,
	},
	{
		Code:  "B",
		Stock: 10,
		X:     19,
		Y:     17,
	},
	{
		Code:  "C",
		Stock: 10,
		X:     10,
		Y:     18,
	},
	{
		Code:  "D",
		Stock: 10,
		X:     21,
		Y:     14,
	},
	{
		Code:  "E",
		Stock: 10,
		X:     8,
		Y:     10,
	},
	{
		Code:  "F",
		Stock: 10,
		X:     16,
		Y:     9,
	},
	{
		Code:  "G",
		Stock: 10,
		X:     11,
		Y:     5,
	},
	{
		Code:  "H",
		Stock: 10,
		X:     17,
		Y:     4,
	},
}

type WarehouseRepository struct {
}

func NewWarehouseRepository() repositories.WarehouseRepository {
	return &WarehouseRepository{}
}

func (r *WarehouseRepository) GetAll() ([]models.Warehouse, error) {
	return WAREHOUSES, nil
}
