package services

import (
	"errors"
	"math"

	"github.com/ahmadkarlam/go-shipping/common/constants"
	"github.com/ahmadkarlam/go-shipping/common/helpers"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/dto"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/models"
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

	id, distance := s.findWarehouse(warehouses, location)

	warehouse, err := s.repository.FindById(id)
	if err != nil {
		return dto.SendingCost{}, errors.New("warehouse not found")
	}
	if warehouse.Stock == 0 {
		return dto.SendingCost{}, errors.New("out of stock")
	}
	if err := s.repository.DecreaseStock(&warehouse); err != nil {
		return dto.SendingCost{}, err
	}

	var warehouseDto dto.Warehouse
	if err := copier.Copy(&warehouseDto, &warehouse); err != nil {
		return dto.SendingCost{}, err
	}

	cost := distance * constants.WarehouseDeliveryPrice
	days := distance / 8
	hours := distance % 8
	sendingCost := dto.SendingCost{
		From:     warehouseDto,
		Distance: distance,
		Cost:     cost,
		Day:      days,
		Hour:     hours,
	}

	return sendingCost, nil
}

func (s *warehouseService) findWarehouse(warehouses []models.Warehouse, location dto.SendVaccineToLocationRequest) (uint, int) {
	distance := math.MaxInt8
	var id uint
	for _, warehouse := range warehouses {
		if warehouse.Stock == 0 {
			continue
		}

		total := helpers.Abs(location.X-warehouse.X) + helpers.Abs(location.Y-warehouse.Y)
		if total < distance {
			distance = total
			id = warehouse.ID
		}
	}

	return id, distance
}
