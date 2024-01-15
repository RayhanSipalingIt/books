package main

import (
	"fmt"

	"github.com/RayhanSipalingIt/novel/controllers/productcontroller"
	"github.com/RayhanSipalingIt/novel/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	// r.GET("/api/products/:id", productcontroller.Show)
	r.GET("/api/products/:id", productcontroller.Index)
	r.GET("/api/products/tipe_buku/:tipe_buku", productcontroller.Index)
	r.GET("/api/products/genre/:genre", productcontroller.Index)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products/:id", productcontroller.Delete)

	// Ubah port menjadi 8081
	port := 8081
	r.Run(fmt.Sprintf(":%d", port))
}
