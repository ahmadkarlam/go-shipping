package mysql

import (
	"github.com/jinzhu/gorm"

	"github.com/ahmadkarlam/go-shipping/modules/warehouses/models"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/repositories"
)

type WarehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) repositories.WarehouseRepository {
	return &WarehouseRepository{
		db: db,
	}
}

func (r *WarehouseRepository) GetAll() ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	err := r.db.Find(&warehouses).Error

	return warehouses, err
}

func (r *WarehouseRepository) DecreaseStock(warehouse models.Warehouse) error {
	return r.db.Model(&warehouse).
		UpdateColumn("stock", gorm.Expr("stock - ?", 1)).
		Error
}
