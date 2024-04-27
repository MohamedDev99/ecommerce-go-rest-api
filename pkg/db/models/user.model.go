package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// user model
type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Email        string             `json:"email" bson:"email,omitempty" validate:"required,email,unique"`
	Password     string             `json:"password" bson:"password,omitempty" validate:"required"`
	FirstName    string             `json:"firstName" bson:"firstName,omitempty" validate:"required,alphanum,min=2,max=20"`
	LastName     string             `json:"lastName" bson:"lastName,omitempty" validate:"required,alphanum,min=2,max=20"`
	Role         string             `json:"role" bson:"role,omitempty" validate:"required,oneof=admin user"`
	Phone        string             `json:"phone" bson:"phone,omitempty" validate:"required,numeric,min=10,max=10"`
	Token        string             `json:"token" bson:"token" `
	RefreshToken string             `json:"refreshToken" bson:"refreshToken"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	UserCart     []ProductUser      `json:"userCart" bson:"userCart"`
	Addresses    []Address          `json:"addresses" bson:"addresses"`
	Orders       []Order            `json:"orders" bson:"orders"`
}

// product user model
type ProductUser struct {
	ProductId   primitive.ObjectID `json:"productId" bson:"productId"`
	ProductName string             `json:"productName" bson:"productName" validate:"required"`
	Price       uint64             `json:"price" bson:"price" validate:"required"`
	Rating      uint8              `json:"rating" bson:"rating"`
	Image       string             `json:"image" bson:"image"`
}

// address model
type Address struct {
	AddressID primitive.ObjectID `json:"addressId" bson:"addressId"`
	House     string             `json:"house" `
	Street    string             `json:"street"`
	City      string             `json:"city"`
	ZipCode   string             `json:"zipCode"`
}
