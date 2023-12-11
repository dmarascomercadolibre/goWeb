package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Item represents a product item with various attributes.
type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// jsonData is a byte slice containing a JSON array of product data.
var jsonData = []byte(`
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

// parseJSON parses the given JSON data and returns a slice of Item structs.
func parseJSON(jsonData []byte) ([]Item, error) {
	var items []Item
	err := json.Unmarshal(jsonData, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// GetAllProducts retrieves all products and returns them as a JSON response.
// The retrieved items are then returned as an HTTP OK response.
func GetAllProducts(c *gin.Context) {
	items, err := parseJSON(jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetProductById retrieves a product by its ID from the provided JSON data.
// If not found, it returns a JSON response with HTTP status code 404 and a "Product not found" message.
func GetProductById(c *gin.Context, id string) {
	items, err := parseJSON(jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	for _, item := range items {
		if item.ID == idInt {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

// GetProductsByPrice retrieves a list of products from the given JSON data
// that have a price higher than the specified price.
func GetProductsByPrice(c *gin.Context, price string) {
	items, err := parseJSON(jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return
	}
	var itemsByPrice []Item
	for _, item := range items {
		if item.Price > priceFloat {
			itemsByPrice = append(itemsByPrice, item)
		}
	}
	c.JSON(http.StatusOK, itemsByPrice)
}

func main() {
	server := gin.Default()

	// Ping test
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!!")
	})

	products := server.Group("/products")

	// Get all products
	products.GET("/", func(c *gin.Context) {
		GetAllProducts(c)
	})
	// Get product by id
	products.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		GetProductById(c, id)
	})

	// get method to get all products that the price is greater than the price passed as a parameter
	products.GET("/price/:price", func(c *gin.Context) {
		price := c.Param("price")
		GetProductsByPrice(c, price)
	})

	server.Run(":8080")
}
