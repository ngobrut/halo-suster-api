package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret  string
	BcryptSalt string
	Postgres   Postgres
	AWS        AWS
}

type Postgres struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	Params   string
}

type AWS struct {
	ID        string
	SecretKey string
	Bucket    string
	Region    string
}

func New() Config {
	_ = godotenv.Load()
	return Config{
		JWTSecret:  os.Getenv("JWT_SECRET"),
		BcryptSalt: os.Getenv("BCRYPT_SALT"),
		Postgres: Postgres{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Params:   os.Getenv("DB_PARAMS"),
		},
		AWS: AWS{
			ID:        os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			Bucket:    os.Getenv("AWS_S3_BUCKET_NAME"),
			Region:    os.Getenv("AWS_REGION"),
		},
	}
}
