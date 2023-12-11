package main

import (
	"net/http"
	"practice/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	hd := handler.NewDefaultProducts()
	server := gin.Default()

	//Ping test
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!!")
	})

	products := server.Group("/products")

	products.POST("/", hd.CreateProduct())
	products.GET("/", hd.GetAllProducts())
	products.GET("/:id", hd.GetProductById())
	products.GET("/price/:price", hd.GetProductsByPrice())
	products.POST("/range", hd.CreateProducts())

	server.Run(":8080")
}

// jsonData is a byte slice containing a JSON array of product data.
var JsonData = []byte(`
[
	{
		"id": 1,
		"name": "Cheese",
		"quantity": 10,
		"code_value": "ABC123",
		"is_published": true,
		"expiration": "2022-12-31",
		"price": 9.99
	},
	{
		"id": 2,
		"name": "Jam",
		"quantity": 5,
		"code_value": "XYZ789",
		"is_published": false,
		"expiration": "2023-06-30",
		"price": 19.99
	},
	{
		"id": 3,
		"name": "Milk",
		"quantity": 20,
		"code_value": "LMN456",
		"is_published": true,
		"expiration": "2023-01-31",
		"price": 29.99
	},
	{
		"id": 4,
		"name": "Bread",
		"quantity": 15,
		"code_value": "QWE123",
		"is_published": true,
		"expiration": "2023-03-31",
		"price": 39.99
	},
	{
		"id": 5,
		"name": "Butter",
		"quantity": 10,
		"code_value": "ASD456",
		"is_published": true,
		"expiration": "2023-04-30",
		"price": 49.99
	},
	{
		"id": 6,
		"name": "Eggs",
		"quantity": 12,
		"code_value": "ZXC789",
		"is_published": true,
		"expiration": "2023-05-31",
		"price": 59.99
	}
]
`)
