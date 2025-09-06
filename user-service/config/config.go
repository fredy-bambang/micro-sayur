package config

import "github.com/spf13/viper"

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`

	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer    string `json:"jwt_issuer"`

	UrlForgotPassword string `json:"url_forgot_password"`
}

type PsqlDB struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	DBName    string `json:"db_name"`
	DBMaxOpen int    `json:"db_max_open"`
	DBMaxIdle int    `json:"db_max_idle"`
}

type RabbitMQ struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	App      App      `json:"app"`
	Psql     PsqlDB   `json:"db"`
	RabbitMQ RabbitMQ `json:"rabbitmq"`
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort:           viper.GetString("APP_PORT"),
			AppEnv:            viper.GetString("APP_ENV"),
			JwtSecretKey:      viper.GetString("JWT_SECRET_KEY"),
			JwtIssuer:         viper.GetString("JWT_ISSUER"),
			UrlForgotPassword: viper.GetString("URL_FORGOT_PASSWORD"),
		},
		Psql: PsqlDB{
			Host:      viper.GetString("DATABASE_HOST"),
			Port:      viper.GetInt("DATABASE_PORT"),
			User:      viper.GetString("DATABASE_USER"),
			Password:  viper.GetString("DATABASE_PASSWORD"),
			DBName:    viper.GetString("DATABASE_NAME"),
			DBMaxOpen: viper.GetInt("DATABASE_MAX_OPEN"),
			DBMaxIdle: viper.GetInt("DATABASE_MAX_IDLE"),
		},
		RabbitMQ: RabbitMQ{
			Host:     viper.GetString("RABBITMQ_HOST"),
			Port:     viper.GetInt("RABBITMQ_PORT"),
			User:     viper.GetString("RABBITMQ_USER"),
			Password: viper.GetString("RABBITMQ_PASSWORD"),
		},
	}
}
