package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/quadratic", quadratic)

	router.Run("0.0.0.0:8080")
}

type request struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
}

func quadratic(c *gin.Context) {
	var request request
	if err := c.BindJSON(&request); err != nil {
		return
	}
	var A = float64(request.A)
	var B = float64(request.B)
	var C = float64(request.C)

	discriminant := math.Pow(B, 2) - 4*A*C

	if discriminant < 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Complex root. Not yet implemented!"})
	} else if discriminant == 0 {
		x := -request.B / (2 * request.A)
		c.IndentedJSON(http.StatusOK, gin.H{"x": x})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"x1": (B + math.Sqrt(discriminant)) / (2*A),
			"x2": (B - math.Sqrt(discriminant)) / (2*A),
		})
	}

}
