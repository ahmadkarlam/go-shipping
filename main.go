package main

import (
	"errors"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type warehouse struct {
	Location coordinate `json:"location"`
	Code     string     `json:"code"`
	Stock    int        `json:"stock"`
}

var WAREHOUSES = []warehouse{
	{
		Code:  "A",
		Stock: 10,
		Location: coordinate{
			X: 5,
			Y: 18,
		},
	},
	{
		Code:  "B",
		Stock: 10,
		Location: coordinate{
			X: 19,
			Y: 17,
		},
	},
	{
		Code:  "C",
		Stock: 10,
		Location: coordinate{
			X: 10,
			Y: 18,
		},
	},
	{
		Code:  "D",
		Stock: 10,
		Location: coordinate{
			X: 21,
			Y: 14,
		},
	},
	{
		Code:  "E",
		Stock: 10,
		Location: coordinate{
			X: 8,
			Y: 10,
		},
	},
	{
		Code:  "F",
		Stock: 10,
		Location: coordinate{
			X: 16,
			Y: 9,
		},
	},
	{
		Code:  "G",
		Stock: 10,
		Location: coordinate{
			X: 11,
			Y: 5,
		},
	},
	{
		Code:  "H",
		Stock: 10,
		Location: coordinate{
			X: 17,
			Y: 4,
		},
	},
}

const PRICE = 5000

func main() {
	if err := serve(); err != nil {
		log.Println(err)
	}
}

func shorterPath(location coordinate) (warehouse, int, error) {
	optimumCost := math.MaxInt8
	index := -1
	for i, warehouse := range WAREHOUSES {
		if warehouse.Stock == 0 {
			continue
		}

		total := abs(location.X-warehouse.Location.X) + abs(location.Y-warehouse.Location.Y)
		if optimumCost > total {
			optimumCost = total
			index = i
		}
	}

	if WAREHOUSES[index].Location.X == 0 {
		return warehouse{}, 0, errors.New("warehouse not found, out of stock")
	}
	WAREHOUSES[index].Stock -= 1

	return WAREHOUSES[index], optimumCost, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func serve() error {
	r := gin.Default()
	r.GET("/warehouse", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": WAREHOUSES,
		})
	})
	r.POST("/get-vaccine", func(ctx *gin.Context) {
		var request coordinate
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		warehouse, cost, err := shorterPath(request)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
			return
		}

		price := cost * PRICE
		days := cost / 8
		hours := cost % 8

		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"warehouse": warehouse,
				"cost":      cost,
				"price":     price,
				"time_took": gin.H{
					"days":  days,
					"hours": hours,
				},
			},
		})
	})
	return r.Run()
}
