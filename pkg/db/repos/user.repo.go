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

type UserRepoAdapter struct {
	mongodbClient *mongo.Client
}

// to have only one instance peer repo
var userCollection *mongo.Collection

// new user repository
func NewUserRepoAdapter(mongodbClient *mongo.Client) *UserRepoAdapter {
	// initialize user collection
	userCollection = db.GetCollection(mongodbClient, "users")
	// initialize user repository
	return &UserRepoAdapter{
		mongodbClient: mongodbClient,
	}
}

// create user
func (u *UserRepoAdapter) CreateUser(user models.User) error {
	// context of mongodb
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// insert user
	_, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		return errors.New(responses.ErrWhileInsertingUser + ": " + err.Error())
	}

	return nil
}

// get all users
func (u *UserRepoAdapter) FindAllUsers() ([]models.User, error) {
	// context of mongodb
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// get all users
	var users []models.User

	cursor, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, errors.New(responses.ErrWhileGettingUsers + ": " + err.Error())
	}

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

// get user by id
func (u *UserRepoAdapter) FindUserByUserId(id string) (models.User, error) {
	// context of mongodb
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// get user by id
	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return models.User{}, errors.New(responses.ErrUserNotFound + ": " + err.Error())
	}

	return user, nil
}

// update user
func (u *UserRepoAdapter) UpdateUser(user models.User) error {
	// context of mongodb
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// update user
	update, err := userCollection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})

	if err != nil {
		return errors.New(responses.ErrWhileUpdatingUser + ": " + err.Error())
	}

	if update.ModifiedCount == 0 {
		return errors.New(responses.ErrUserNotFound)
	}

	return nil
}

// delete user
func (u *UserRepoAdapter) DeleteUser(id string) error {
	// context of mongodb
	ctx, ctcErr := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctcErr()

	// delete user
	delete, err := userCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return errors.New(responses.ErrWhileDeletingUser + ": " + err.Error())
	}

	if delete.DeletedCount == 0 {
		return errors.New(responses.ErrUserNotFound)
	}

	return nil
}
