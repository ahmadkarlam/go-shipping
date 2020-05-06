package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"

	"github.com/ahmadkarlam/go-shipping/common/constants"
	mockRepositories "github.com/ahmadkarlam/go-shipping/mocks/modules/warehouses/repositories"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/dto"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/models"
)

func Test_warehouseService_FindNearbyWarehouse(t *testing.T) {
	type args struct {
		location dto.SendVaccineToLocationRequest
	}
	tests := []struct {
		name          string
		args          args
		configureMock func(repository *mockRepositories.MockWarehouseRepository, args args)
		want          dto.SendingCost
		wantErr       bool
	}{
		{
			name: "Success find nearby warehouse",
			args: args{
				location: dto.SendVaccineToLocationRequest{
					X: 1,
					Y: 1,
				},
			},
			configureMock: func(repository *mockRepositories.MockWarehouseRepository, args args) {
				warehouses := []models.Warehouse{
					{
						Model: gorm.Model{ID: 1},
						Code:  "A",
						Stock: 10,
						X:     2,
						Y:     2,
					},
					{
						Model: gorm.Model{ID: 1},
						Code:  "B",
						Stock: 10,
						X:     3,
						Y:     3,
					},
				}
				repository.EXPECT().GetAll().Return(warehouses, nil)
				warehouse := models.Warehouse{
					Model: gorm.Model{ID: 1},
					Code:  "A",
					Stock: 10,
					X:     2,
					Y:     2,
				}
				repository.EXPECT().FindById(warehouses[0].ID).Return(warehouse, nil)
				repository.EXPECT().DecreaseStock(warehouse).DoAndReturn(func(models.Warehouse) (models.Warehouse, error) {
					warehouse.Stock -= 1
					return warehouse, nil
				})
			},
			want: dto.SendingCost{
				From: dto.Warehouse{
					Code:  "A",
					Stock: 9,
					X:     2,
					Y:     2,
				},
				Distance: 2,
				Cost:     2 * constants.WarehouseDeliveryPrice,
				Day:      0,
				Hour:     2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mockRepositories.NewMockWarehouseRepository(ctrl)
			tt.configureMock(repository, tt.args)
			s := &warehouseService{
				repository: repository,
			}
			got, err := s.FindNearbyWarehouse(tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindNearbyWarehouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindNearbyWarehouse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_warehouseService_GetAllWarehouse(t *testing.T) {
	tests := []struct {
		name          string
		configureMock func(repository *mockRepositories.MockWarehouseRepository)
		want          []dto.Warehouse
		wantErr       bool
	}{
		{
			name: "success get all warehouse",
			configureMock: func(repository *mockRepositories.MockWarehouseRepository) {
				repository.EXPECT().GetAll().Return([]models.Warehouse{
					{Code: "A"},
				}, nil)
			},
			want: []dto.Warehouse{
				{Code: "A"},
			},
			wantErr: false,
		},
		{
			name: "error database when fetch all warehouses",
			configureMock: func(repository *mockRepositories.MockWarehouseRepository) {
				repository.EXPECT().GetAll().Return([]models.Warehouse{}, errors.New("database error"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mockRepositories.NewMockWarehouseRepository(ctrl)
			s := &warehouseService{
				repository: repository,
			}
			tt.configureMock(repository)
			got, err := s.GetAllWarehouse()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllWarehouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllWarehouse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
