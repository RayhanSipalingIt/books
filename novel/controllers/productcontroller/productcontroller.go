package productcontroller

import (
	"net/http"

	"github.com/RayhanSipalingIt/novel/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var products []models.Product

	// Ambil nilai dari parameter pencarian (id, tipe_buku, genre)
	id := c.Param("id")
	tipeBuku := c.Param("tipe_buku")
	genre := c.Param("genre")

	// Mengecek apakah ada parameter pencarian dan menjalankan pencarian berdasarkan satu parameter saja
	if id != "" {
		models.DB.Where("id = ?", id).Find(&products)
	} else if tipeBuku != "" {
		models.DB.Where("tipe_buku = ?", tipeBuku).Find(&products)
	} else if genre != "" {
		models.DB.Where("genre = ?", genre).Find(&products)
	} else {
		// Jika tidak ada parameter pencarian, ambil semua data
		models.DB.Find(&products)
	}

	// Handle jika hasil pencarian kosong
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {

}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
