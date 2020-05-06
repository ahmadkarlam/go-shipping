package mysql

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"

	"github.com/ahmadkarlam/go-shipping/infrastructure/database"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/models"
)

func TestWarehouseRepository_GetAll(t *testing.T) {
	query := "SELECT \\* FROM `warehouses` (.+)"
	tests := []struct {
		name          string
		configureMock func(mock sqlmock.Sqlmock)
		want          []models.Warehouse
		wantErr       bool
	}{
		{
			name: "Success get all warehouse",
			configureMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"code"}).
					AddRow("A").
					AddRow("B")
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []models.Warehouse{
				{Code: "A"},
				{Code: "B"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := database.DBTest()

			tt.configureMock(mock)
			r := &WarehouseRepository{db: db}
			got, err := r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestWarehouseRepository_DecreaseStock(t *testing.T) {
	type args struct {
		warehouse *models.Warehouse
	}
	query := "UPDATE `warehouses` (.+)"
	tests := []struct {
		name          string
		args          args
		configureMock func(mock sqlmock.Sqlmock, args args)
		wantErr       bool
	}{
		{
			name: "Success decrease stock vaccine",
			args: args{warehouse: &models.Warehouse{
				Model: gorm.Model{ID: 1},
				Stock: 10,
			}},
			configureMock: func(mock sqlmock.Sqlmock, args args) {
				mock.ExpectBegin()
				mock.ExpectExec(query).WithArgs(1, args.warehouse.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := database.DBTest()

			tt.configureMock(mock, tt.args)
			r := &WarehouseRepository{
				db: db,
			}
			if err := r.DecreaseStock(tt.args.warehouse); (err != nil) != tt.wantErr {
				t.Errorf("DecreaseStock() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestWarehouseRepository_FindById(t *testing.T) {
	type args struct {
		id uint
	}

	query := "SELECT \\* FROM `warehouses`  WHERE (.+)"
	tests := []struct {
		name          string
		args          args
		configureMock func(mock sqlmock.Sqlmock, args args)
		want          models.Warehouse
		wantErr       bool
	}{
		{
			name: "Success find warehouse by id",
			args: args{id: 1},
			configureMock: func(mock sqlmock.Sqlmock, args args) {
				rows := sqlmock.NewRows([]string{"id", "code", "stock"}).AddRow(1, "A", 10)
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: models.Warehouse{
				Model: gorm.Model{ID: 1},
				Code:  "A",
				Stock: 10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := database.DBTest()

			tt.configureMock(mock, tt.args)
			r := &WarehouseRepository{
				db: db,
			}
			got, err := r.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
