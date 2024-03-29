package middleware

import (
	"net/http"

	auth "hotelbooking-go/Auth"
	Init "hotelbooking-go/initializer"
	"hotelbooking-go/models"

	"github.com/gin-gonic/gin"
)

// AuthMiddleWare for OwnerVerificaton
func AuthMiddleWare(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		c.JSON(400, gin.H{"error": "token missing"})
		c.Abort()
		return
	}
	rslt, err := auth.Trim(header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "trim error"})
		c.Abort()
		return
	}
	var owner models.Owner
	result := Init.DB.Where("user_name = ?", rslt).First(&owner)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username not found"})
		c.Abort()
		return
	}

	c.Next()
}
