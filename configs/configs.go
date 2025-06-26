package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type DBC struct {
	Driver   string
	Type     string
	Username string
	Password string
	Hostname string
	Port     string
	Database string
	Sslmode  string
}

type JWT struct {
	SecretKey []byte
}

type Mail struct {
	Server   string
	Port     string
	Email    string
	Password string
}

type Config struct {
	DBC  DBC
	JWT  JWT
	Mail Mail
}

func LoadConfigs() Config {
	godotenv.Load()
	return Config{
		DBC{
			os.Getenv("DB_DRIVER"),
			os.Getenv("DB_TYPE"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PWD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSL"),
		},
		JWT{
			[]byte(os.Getenv("JWT_SECRET")),
		}, Mail{
			os.Getenv("EMAIL_HOST"),
			os.Getenv("EMAIL_PORT"),
			os.Getenv("EMAIL_HOST_USER"),
			os.Getenv("EMAIL_HOST_PASSWORD"),
		}}
}
