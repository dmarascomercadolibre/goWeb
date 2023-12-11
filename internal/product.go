package internal

// Item represents a product item with various attributes.
type Product struct {
	ID int `json:"id"`
	AtributtesProduct
}

// AtributtesItem represents various attributes of an Item.
type AtributtesProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}
