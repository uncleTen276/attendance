package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Host string `envconfig:"APP_HOST" default:"localhost"`
		PORT string `envconfig:"APP_PORT" default:":8080"`
	}
	Database struct {
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Name     string `envconfig:"DB_NAME" default:"greendeco"`
		Port     string `envconfig:"DB_PORT" default:"5432"`
		User     string `envconfig:"DB_USER" default:"greendeco"`
		Password string `envconfig:"DB_PASSWORD" default:"greedenco"`
		SSLMode  string `envconfig:"SSL_MODE" default:"require"`
	}
	Auth struct {
		JWTSecret        string `envconfig:"JWT_SECRET" default:"token-secret"`
		TokenExpire      int    `envconfig:"TOKEN_EXPIRE" default:"60"`
		ShortTokenExpire int    `envconfig:"SHORT_TOKEN_EXPIRE" default:"15"`
	}
	Contract struct {
		Infura_Url       string `envconfig:"INFURA_URL" default:"http://127.0.0.1:8545/"`
		Contract_Address string `envconfig:"CONTRACT_ADDRESS" default:"0x5FbDB2315678afecb367f032d93F642f64180aa3"`
		Private_Key      string `envconfig:"PRIVATE_KEY" default:"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"`
	}
}

var appConfig = &Config{}

func AppConfig() *Config {
	return appConfig
}

func LoadConfig() error {
	if err := envconfig.Process("", appConfig); err != nil {
		return err
	}

	return nil
}
