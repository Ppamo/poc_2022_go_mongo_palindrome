package entities

type Product struct {
	ID          int    `json:"id"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int64  `json:"price"`
}
