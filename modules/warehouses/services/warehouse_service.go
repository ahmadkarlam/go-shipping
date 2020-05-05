package services

import (
	"errors"
	"math"

	"github.com/ahmadkarlam/go-shipping/common/constants"
	"github.com/ahmadkarlam/go-shipping/common/helpers"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/dto"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/repositories"

	"github.com/jinzhu/copier"
)

type WarehouseService interface {
	GetAllWarehouse() ([]dto.Warehouse, error)
	FindNearbyWarehouse(location dto.SendVaccineToLocationRequest) (dto.SendingCost, error)
}

type warehouseService struct {
	repository repositories.WarehouseRepository
}

func NewWarehouseService(repository repositories.WarehouseRepository) WarehouseService {
	return &warehouseService{
		repository: repository,
	}
}

func (s *warehouseService) GetAllWarehouse() ([]dto.Warehouse, error) {
	warehouses, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	var warehouse []dto.Warehouse
	if err := copier.Copy(&warehouse, &warehouses); err != nil {
		return nil, err
	}

	return warehouse, nil
}

func (s *warehouseService) FindNearbyWarehouse(location dto.SendVaccineToLocationRequest) (dto.SendingCost, error) {
	warehouses, err := s.repository.GetAll()
	if err != nil {
		return dto.SendingCost{}, nil
	}

	distance := math.MaxInt8
	index := -1
	for i, warehouse := range warehouses {
		if warehouse.Stock == 0 {
			continue
		}

		total := helpers.Abs(location.X-warehouse.X) + helpers.Abs(location.Y-warehouse.Y)
		if distance > total {
			distance = total
			index = i
		}
	}

	if warehouses[index].X == 0 {
		return dto.SendingCost{}, errors.New("warehouse not found, out of stock")
	}
	// TODO: update to database
	warehouses[index].Stock -= 1

	var warehouse dto.Warehouse
	if err := copier.Copy(&warehouse, &warehouses[index]); err != nil {
		return dto.SendingCost{}, err
	}

	cost := distance * constants.WarehouseDeliveryPrice
	days := distance / 8
	hours := distance % 8
	sendingCost := dto.SendingCost{
		From:     warehouse,
		Distance: distance,
		Cost:     cost,
		Day:      days,
		Hour:     hours,
	}

	return sendingCost, nil
}
