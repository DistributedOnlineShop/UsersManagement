package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment         string        `mapstructure:"ENVIRONMENT"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	AllowHeaders        []string      `mapstructure:"ALLOW_HEADERS"`
	RedisAddress        string        `mapstructure:"REDIS_ADDRESS"`
	RedisPassword       string        `mapstructure:"REDIS_PASSWORD"`
	MigrationURL        string        `mapstructure:"MIGRATION_URL"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	TimeCost            uint32        `mapstructure:"TIME_COST"`
	MemoryCost          uint32        `mapstructure:"MEMORY_COST"`
	Parallelism         uint8         `mapstructure:"PARALLELISM"`
	KeyLength           uint32        `mapstructure:"KEY_LENGTH"`
	UsersManagementPort string        `mapstructure:"USERS_MANAGEMENT_PORT"`
	KeySeed             string        `mapstructure:"KEY_SEED"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)

	return config, nil
}
