package types

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartRepository interface {
	AddToCart(userId, productId primitive.ObjectID) error
	RemoveItemFromCart(productId string) error
	GetItemFromCart(product models.Product) (models.Product, error)
	// GetCartItems() ([]models.Product, error)
	BuyFromCart(product models.Product) error
	InstantBuy(product models.Product) error
}
