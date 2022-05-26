package entities

type Product struct {
	ID          int    `bson:"id"`
	Brand       string `bson:"brand"`
	Description string `bson:"description"`
	Image       string `bson:"image"`
	Price       int64  `bson:"price"`
}
