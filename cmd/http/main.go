package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ahmadkarlam/go-shipping/common/resolver"
	"github.com/ahmadkarlam/go-shipping/modules/warehouses/handlers"
)

func main() {
	if err := serve(); err != nil {
		panic(err)
	}
}

func serve() error {
	r := gin.Default()

	resolver := resolver.NewResolver()

	warehouseHandler := handlers.NewWarehouseHandler(resolver)

	r.GET("/warehouse", warehouseHandler.GetAll)
	r.POST("/warehouse/send-vaccine", warehouseHandler.SendVaccineToLocation)
	return r.Run()
}
