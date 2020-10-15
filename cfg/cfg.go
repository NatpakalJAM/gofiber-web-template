package cfg

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//Config main config
type Config struct {
	AppURL string `mapstructure:"app-url"`
	DB     struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Database string `mapstructure:"database"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"db"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Db       int    `mapstructure:"db"`
		Password string `mapstructure:"password"`
	} `mapstructure:"redis"`

	RabbitMQ struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Queues   struct {
			Process string `mapstructure:"process"`
		} `mapstructure:"queues"`
	} `mapstructure:"rabbitmq"`

	Environment string
	Instance    string //blue green
}

//C config instance
var C Config

//Init init cfg
func Init() {
	env := os.Getenv("ENVIRONMENT")
	fmt.Println("Environment " + env)
	cfgName := ""
	if env == "production" {
		cfgName = "config.production"
	} else {
		cfgName = "config.development"
	}
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(cfgName)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
	C.Environment = env
	C.Instance = os.Getenv("INSTANCE")
}
