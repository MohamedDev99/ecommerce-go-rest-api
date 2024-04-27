package controllers

import (
	"time"

	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/responses"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/validation"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs/token"
	types "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs/types/out"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authRepo types.AuthRepository
}

func NewAuthController(authRepo types.AuthRepository) *AuthController {
	// initialize auth controller
	return &AuthController{authRepo: authRepo}
}

// login
func (a *AuthController) Login(c fiber.Ctx) error {
	// validation
	loginBody, err := validation.ValidateLoginBody(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	// login
	user, err := a.authRepo.Login(loginBody.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.AuthFailResponse{Status: responses.FailStatus, Message: responses.ErrAuthInvalidEmailOrPassword})
	}
	//check password
	if ok := comparePassword(user.Password, loginBody.Password); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.AuthFailResponse{Status: responses.FailStatus, Message: responses.ErrAuthInvalidEmailOrPassword})
	}

	// generate token
	token, err := token.GenerateToken(user.Email, user.FirstName, user.LastName, user.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.AuthFailResponse{Status: responses.FailStatus, Message: responses.ErrInternalServerError})
	}

	// pass token
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(responses.AuthSuccessResponseWithUser{Status: responses.SuccessStatus, Message: responses.AuthSuccess, User: user})
}

// register
func (a *AuthController) Register(c fiber.Ctx) error {

	// validation
	registerBody, err := validation.ValidateRegisterBody(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// hash password
	hashedPassword, err := hashPassword(registerBody.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.AuthFailResponse{Status: responses.FailStatus, Message: responses.ErrAuthInvalidEmailOrPassword})
	}

	// generate user id
	userID := primitive.NewObjectID()
	// generate token
	token, err := token.GenerateToken(registerBody.Email, registerBody.FirstName, registerBody.LastName, userID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.AuthFailResponse{Status: responses.FailStatus, Message: responses.ErrInternalServerError})
	}

	// register
	user := models.User{
		Email:        registerBody.Email,
		Password:     hashedPassword,
		FirstName:    registerBody.FirstName,
		LastName:     registerBody.LastName,
		Role:         registerBody.Role,
		Phone:        registerBody.Phone,
		UserId:       userID,
		Token:        token,
		RefreshToken: "",
		UserCart:     []models.ProductUser{},
		Addresses:    []models.Address{},
		Orders:       []models.Order{},
	}
	// insert user
	err = a.authRepo.Register(user)

	if err != nil {
		if e, ok := err.(error); ok {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.AuthFailResponse{Status: responses.FailStatus, Message: responses.ErrUserNotCreated + e.Error()})
		}
	}

	// generate jwt token and pass it to frontend
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(responses.AuthSuccessResponseWithUser{Status: responses.SuccessStatus, Message: responses.UserCreated, User: user})
}

// hash password
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

// compare password
func comparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Info("compare password: ", zap.Any("err", err.Error()))
	}
	return err == nil
}
