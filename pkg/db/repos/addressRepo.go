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

type AddressRepo struct {
	mongodbClient *mongo.Client
}

var addressCollection *mongo.Collection // addressCollection

func NewAddressRepoAdapter(mongodbClient *mongo.Client) *AddressRepo {
	// initialize address collection
	addressCollection = db.GetCollection(mongodbClient, "addresses")
	// initialize address repository
	return &AddressRepo{mongodbClient: mongodbClient}
}

// create address
func (*AddressRepo) AddAddress(address models.Address) error {
	// mongo context
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// insert address
	insertResult, err := addressCollection.InsertOne(ctx, address)

	if err != nil {
		return errors.New(responses.ErrWhileInsertingAddress + ": " + err.Error())
	}

	if insertResult.InsertedID == nil {
		return errors.New(responses.ErrWhileInsertingAddress)
	}

	return nil
}

// edit home address
func (*AddressRepo) EditHomeAddress(addressId, homeAddress string) error {
	// mongo context
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// update address
	updateResult, err := addressCollection.UpdateOne(ctx, bson.M{"_id": addressId}, bson.M{"$set": bson.M{"homeAddress": homeAddress}})

	if err != nil {
		return errors.New(responses.ErrWhileUpdatingAddress + ": " + err.Error())
	}

	if updateResult.MatchedCount == 0 {
		return errors.New(responses.ErrAddressNotFound)
	}

	if updateResult.ModifiedCount == 0 {
		return errors.New(responses.ErrWhileUpdatingAddress)
	}

	return nil
}

// edit work address
func (*AddressRepo) EditWorkAddress(addressId, workAddress string) error {
	// mongo context
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// update address
	updateResult, err := addressCollection.UpdateOne(ctx, bson.M{"_id": addressId}, bson.M{"$set": bson.M{"workAddress": workAddress}})

	if err != nil {
		return errors.New(responses.ErrWhileUpdatingAddress + ": " + err.Error())
	}

	if updateResult.MatchedCount == 0 {
		return errors.New(responses.ErrAddressNotFound)
	}

	if updateResult.ModifiedCount == 0 {
		return errors.New(responses.ErrWhileUpdatingAddress)
	}

	return nil
}

// delete address
func (*AddressRepo) DeleteAddress(addressId string) error {
	// mongo context
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// delete address
	deleteResult, err := addressCollection.DeleteOne(ctx, bson.M{"_id": addressId})

	if err != nil {
		return errors.New(responses.ErrWhileDeletingAddress + ": " + err.Error())
	}

	if deleteResult.DeletedCount == 0 {
		return errors.New(responses.ErrAddressNotFound)
	}

	return nil
}
