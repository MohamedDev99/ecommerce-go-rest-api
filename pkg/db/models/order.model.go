package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// order model
type Order struct {
	OrderID        primitive.ObjectID `json:"orderId" bson:"orderId"`
	OrderCart      []ProductUser      `json:"orderCart" bson:"orderCart"`
	OrderCreatedAt time.Time          `json:"orderCreatedAt"`
	Price          uint64             `json:"price"`
	Discount       uint64             `json:"discount"`
	PaymentMethod  Payment            `json:"paymentMethod" bson:"paymentMethod"`
}

// payment model
type Payment struct {
	Digital bool
	COD     bool
}
