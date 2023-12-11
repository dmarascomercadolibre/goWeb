package handler

import (
	"fmt"
	"net/http"
	"practice/internal"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewDefaultProducts() *DefaultProducts {
	return &DefaultProducts{
		db:     make([]internal.Product, 0),
		lastID: 0,
	}
}

type DefaultProducts struct {
	db     []internal.Product
	lastID int
}

type ProductJSON struct {
	ID int `json:"id"`
	internal.AtributtesProduct
}

type BodyRequestCreate struct {
	internal.AtributtesProduct
}

func (p *DefaultProducts) CreateProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, map[string]any{"message": "invalid token"})
			return
		}
		var body BodyRequestCreate
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		product := internal.Product{
			AtributtesProduct: body.AtributtesProduct,
		}
		p.lastID++
		product.ID = p.lastID

		p.db = append(p.db, product)

		//response
		ctx.JSON(http.StatusCreated, gin.H{"message": "Product created successfully",
			"data": ProductJSON{
				ID:                product.ID,
				AtributtesProduct: product.AtributtesProduct,
			},
		})
	}
}

// function that create a list of products an return it as a JSON response.
func (p *DefaultProducts) CreateProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, map[string]any{"message": "invalid token"})
			return
		}
		var body []BodyRequestCreate
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		for _, product := range body {
			product := internal.Product{
				AtributtesProduct: product.AtributtesProduct,
			}
			p.lastID++
			product.ID = p.lastID

			p.db = append(p.db, product)
		}

		//response
		ctx.JSON(http.StatusCreated, gin.H{"message": "Products created successfully",
			"data": p.db,
		})
	}
}

// GetAllProducts retrieves all products and returns them as a JSON response.
// The retrieved items are then returned as an HTTP OK response.
func (p *DefaultProducts) GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, p.db)
	}
}

// GetProductById retrieves a product by its ID from the provided JSON data.
// If not found, it returns a JSON response with HTTP status code 404 and a "Product not found" message.
func (p *DefaultProducts) GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		for _, product := range p.db {
			if product.ID == idInt {
				ctx.JSON(http.StatusOK, product)
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
	}
}

// GetProductsByPrice retrieves a list of products from the given JSON data
// that have a price higher than the specified price.
func (p *DefaultProducts) GetProductsByPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		price := ctx.Param("price")
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Error converting string to float:", err)
			return
		}
		var productsByPrice []internal.Product
		for _, product := range p.db {
			if product.Price > priceFloat {
				productsByPrice = append(productsByPrice, product)
			}
		}
		ctx.JSON(http.StatusOK, productsByPrice)
	}
}

// // parseJSON parses the given JSON data and returns a slice of Item structs.
// func parseJSON(jsonData []byte) ([]internal.Product, error) {
// 	var products []internal.Product
// 	err := json.Unmarshal(jsonData, &products)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return products, nil
// }
