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
	"go.uber.org/zap"
)

type AuthRepoAdapter struct {
	mongodbClient *mongo.Client
}

// to have only one instance peer repo
var authCollection *mongo.Collection

// auth repository
func NewAuthRepoAdapter(mongodbClient *mongo.Client) *AuthRepoAdapter {
	// initialize auth collection
	authCollection = db.GetCollection(mongodbClient, "users")
	// initialize auth repository
	return &AuthRepoAdapter{mongodbClient: mongodbClient}
}

// login
func (a *AuthRepoAdapter) Login(email string) (models.User, error) {
	// request context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	// search user by email
	filter := bson.M{"email": email}

	err := authCollection.FindOne(ctx, filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		log.Info("User not found :" + err.Error())
		return models.User{}, errors.New("User not found" + err.Error())
	}

	// marshal user
	// userByte, err := bson.Marshal(user)
	// log.Info("userByte : ", zap.Any("userByte", userByte))
	// // unmarshal user
	// _ = bson.Unmarshal(userByte, &user)

	return user, nil
}

// sign up
func (a *AuthRepoAdapter) Register(user models.User) error {
	// request context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check if user already exists
	count, err := authCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Info(responses.ErrWhileCountingUser + ": " + err.Error())
		return errors.New(responses.ErrInternalServerError)
	}

	if count > 0 {
		log.Info(responses.ErrUserAlreadyExists)
		return errors.New(responses.ErrUserAlreadyExists)
	}

	// insert user
	insertResult, err := authCollection.InsertOne(ctx, user)

	if err != nil {
		log.Info(responses.ErrWhileInsertingUser + ": " + err.Error())
		return errors.New(responses.ErrWhileInsertingUser)
	}

	log.Info("insertResult : ", zap.Any("insertResult", insertResult))

	return nil

}
