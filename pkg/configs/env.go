package configs

import (
	"os"

	"github.com/joho/godotenv"
)

// load .env file
func LoadEnvFromDotEnv() {
	error := godotenv.Load()
	if error != nil {
		Logger.Error("Error loading .env file")
	}
}

// get MongoDB URI from .env file
func EnvMongoURI() string {
	// err := loadEnvFromDotEnv()
	return os.Getenv("MONGO_DB_URI")
}

// get JWT Secret from .env file
func EnvJWTSecret() string {
	// err := loadEnvFromDotEnv()
	return os.Getenv("JWT_SECRET")
}

// server PORT
func EnvPort() string {
	// err := loadEnvFromDotEnv()
	return os.Getenv("PORT")
}

// get DB name
func EnvMongoDBName() string {
	// err := loadEnvFromDotEnv()
	return os.Getenv("MONGO_DB_NAME")
}

// api version
func EnvAPIVersion() string {
	return os.Getenv("API_VERSION")
}
