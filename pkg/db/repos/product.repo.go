package repos

import (
	"context"
	"errors"
	"time"

	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/responses"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepo struct {
	mongodbClient *mongo.Client
}

// to have only one instance peer repo
var productCollection *mongo.Collection

func NewProductRepoAdapter(mongodbClient *mongo.Client) *ProductRepo {
	// initialize product collection
	productCollection = db.GetCollection(mongodbClient, "products")
	// initialize product repository
	return &ProductRepo{mongodbClient: mongodbClient}
}

// create product
func (*ProductRepo) CreateProduct(product models.Product) (models.Product, error) {
	// mongo context
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// insert product
	_, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		return models.Product{}, errors.New(responses.ErrWhileInsertingProduct + ": " + err.Error())
	}
	return product, nil
}

// get products (products viewer by admin, ...)

// get product (product viewer by admin and user)
// product viewer by admin
func (productRepo *ProductRepo) ProductViewerByAdmin(id string) (models.Product, error) {
	// mongo context
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// get product
	var product models.Product

	err := productCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return models.Product{}, errors.New(responses.ErrWhileGettingProduct + ": " + err.Error())
	}

	return product, nil
}

// search product for user
func (productRepo *ProductRepo) SearchProduct() ([]models.Product, error) {
	// mongo context
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// get products
	var products []models.Product
	cursor, err := productCollection.Find(ctx, bson.M{})

	if err != nil {
		return []models.Product{}, errors.New(responses.ErrWhileGettingProducts + ": " + err.Error())
	}

	for cursor.Next(ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return []models.Product{}, errors.New(responses.ErrWhileGettingProducts + ": " + err.Error())
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return []models.Product{}, errors.New(responses.ErrWhileGettingProducts + ": " + err.Error())
	}

	return []models.Product{}, nil
}

// search product by query (filter)
func (productRepo *ProductRepo) SearchProductByQuery(query string) ([]models.Product, error) {
	// mongo context
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// get products
	var searchProducts []models.Product
	cursor, err := productCollection.Find(ctx, bson.M{"$text": bson.M{"$search": query}})

	if err != nil {
		return []models.Product{}, errors.New(responses.ErrWhileGettingProducts + ": " + err.Error())
	}

	for cursor.Next(ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return []models.Product{}, errors.New(responses.ErrWhileGettingProducts + ": " + err.Error())
		}
		searchProducts = append(searchProducts, product)
	}

	if err := cursor.Err(); err != nil {
		return []models.Product{}, errors.New(responses.ErrWhileGettingProducts + ": " + err.Error())
	}

	return searchProducts, nil
}

// update product

// delete product
func (*ProductRepo) DeleteProduct(id string) error {
	// mongo context
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// delete product
	deleteProduct, err := productCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return errors.New(responses.ErrWhileDeletingProduct + ": " + err.Error())
	}

	if deleteProduct.DeletedCount == 0 {
		return errors.New(responses.ErrProductNotFound)
	}

	return nil
}
