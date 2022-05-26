package repositories

import (
	"context"
	"entities"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
)

type repo struct{}

func NewMongoRepository() ProductsRepository {
	var err error
	s := fmt.Sprintf("mongodb://%s:%s@%s:27017/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
	log.Printf("> Creating connection to mongo using %s", s)
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(s))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection = client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION"))
	return &repo{}
}

func fillProducts(results []bson.M) ([]entities.Product, error) {
	var products []entities.Product
	for _, item := range results {
		product := entities.Product{}
		b, err := bson.Marshal(item)
		if err != nil {
			return nil, err
		}
		bson.Unmarshal(b, &product)
		products = append(products, product)
	}
	return products, nil
}

func (*repo) FindText(s string) ([]entities.Product, error) {
	query := bson.M{"$or": []interface{}{bson.M{"brand": bson.M{"$regex": s}}, bson.M{"description": bson.M{"$regex": s}}}}
	c, err := collection.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = c.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return fillProducts(results)
}

func (*repo) Find(id int) (entities.Product, error) {
	product := entities.Product{}
	if err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&product); err != nil {
		log.Fatal(err)
	}
	return product, nil
}

func (*repo) FindAll() ([]entities.Product, error) {
	c, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = c.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return fillProducts(results)
}
