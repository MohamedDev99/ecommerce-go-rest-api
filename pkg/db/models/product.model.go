package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// product model
type Product struct {
	ProductID   primitive.ObjectID `json:"productId" bson:"productId"`
	ProductName string             `json:"productName" `
	Price       uint64             `json:"price"`
	Rating      uint8              `json:"rating"`
	Image       string             `json:"image"`
	// will be added later
}
