package entities

type Product struct {
	ID          int    `json:"id" bson:"id"`
	Brand       string `json:"brand" bson:"brand"`
	Description string `json:"description" bson:"description"`
	Image       string `json:"image" bson:"image"`
	Price       int64  `json:"price" bson:"price"`
	Discount    int    `json:"discount"`
}
