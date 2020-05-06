package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ahmadkarlam/go-shipping/common/resolver"
	_ "github.com/ahmadkarlam/go-shipping/docs"
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
	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r.Run()
}
