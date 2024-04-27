package repos

import (
	"context"
	"errors"
	"time"

	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/responses"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepo struct {
	mongodbClient *mongo.Client
}

var cartCollections map[string]*mongo.Collection

// new cart repository
func NewCartRepoAdapter(mongodbClient *mongo.Client) *CartRepo {
	// initialize cart collection
	cartCollections = map[string]*mongo.Collection{
		"users":    db.GetCollection(mongodbClient, "users"),
		"products": db.GetCollection(mongodbClient, "products"),
	}
	// initialize cart repository
	return &CartRepo{mongodbClient: mongodbClient}
}

// add to cart
func (*CartRepo) AddToCart(userId, productId primitive.ObjectID) error {
	// context of mongodb
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// insert cart
	_, err := cartCollections["users"].InsertOne(ctx, models.Cart{}) // TODO : Add cart items

	if err != nil {
		return errors.New(responses.ErrWhileInsertingCart + ": " + err.Error())
	}

	return nil
}

// remove item from cart
func (*CartRepo) RemoveItemFromCart(productId string) error {
	// context of mongodb
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// delete cart
	deleteResult, err := cartCollections["users"].DeleteOne(ctx, bson.M{"products._id": productId})

	if err != nil {
		return errors.New(responses.ErrWhileDeletingCart + ": " + err.Error())
	}

	if deleteResult.DeletedCount == 0 {
		return errors.New(responses.ErrCartNotFound)
	}

	return nil
}

// get item from cart
func (*CartRepo) GetItemFromCart(product models.Product) (models.Product, error) {
	// context of mongodb
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// get cart
	cart := models.Cart{}

	err := cartCollections["users"].FindOne(ctx, bson.M{"products._id": product.ProductID}).Decode(&cart)

	if err != nil {
		return models.Product{}, errors.New(responses.ErrWhileGettingCart + ": " + err.Error())
	}

	if len(cart.Products) == 0 {
		return models.Product{}, errors.New(responses.ErrCartNotFound)
	}

	return models.Product{}, nil
}

// buy from cart
func (*CartRepo) BuyFromCart(product models.Product) error {
	// context of mongodb
	// ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	// defer ctcErr()

	return nil
}

// instant buy
func (*CartRepo) InstantBuy(product models.Product) error {
	// context of mongodb
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	return nil
}
