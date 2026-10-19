package config

import "github.com/spf13/viper"

type Config struct {
	AppName            string
	AppPort            string
	AppEnv             string
	PostgresHost       string
	PostgresPort       string
	PostgresUser       string
	PostgresPassword   string
	PostgresDB         string
	RedisHost          string
	RedisPort          string
	RedisPassword      string
	MinioEndpoint      string
	MinioAccessKey     string
	MinioSecretKey     string
	MinioUseSSL        bool
	MinioBucket        string
	GRPCHost           string
	GRPCPort           string
	GRPCServerPort     string
	CORSAllowedOrigins string
}

func Load() (*Config, error) {
	viper.SetConfigName(".env")

	viper.SetConfigType("env")

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	cfg := &Config{
		AppName:            viper.GetString("APP_NAME"),
		AppPort:            viper.GetString("APP_PORT"),
		AppEnv:             viper.GetString("APP_ENV"),
		PostgresHost:       viper.GetString("POSTGRES_HOST"),
		PostgresPort:       viper.GetString("POSTGRES_PORT"),
		PostgresUser:       viper.GetString("POSTGRES_USER"),
		PostgresPassword:   viper.GetString("POSTGRES_PASSWORD"),
		PostgresDB:         viper.GetString("POSTGRES_DB"),
		RedisHost:          viper.GetString("REDIS_HOST"),
		RedisPort:          viper.GetString("REDIS_PORT"),
		RedisPassword:      viper.GetString("REDIS_PASSWORD"),
		MinioEndpoint:      viper.GetString("MINIO_ENDPOINT"),
		MinioAccessKey:     viper.GetString("MINIO_ACCESS_KEY"),
		MinioSecretKey:     viper.GetString("MINIO_SECRET_KEY"),
		MinioUseSSL:        viper.GetBool("MINIO_USE_SSL"),
		MinioBucket:        viper.GetString("MINIO_BUCKET"),
		GRPCHost:           viper.GetString("GRPC_HOST"),
		GRPCPort:           viper.GetString("GRPC_PORT"),
		GRPCServerPort:     viper.GetString("GRPC_SERVER_PORT"),
		CORSAllowedOrigins: viper.GetString("CORS_ALLOWED_ORIGINS"),
	}

	if cfg.AppPort == "" {
		cfg.AppPort = "9090"
	}

	return cfg, nil
}
