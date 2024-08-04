package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//#Para usar la ruta local usar esto :    "../config" en CONFIG_PATH  y usar "/root" para crear el dokerfile
	viper.AddConfigPath("/root")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error leyendo archivo de configuración: %s", err)
	}
}

func GetSecret() string {
	return viper.GetString("JWT_SECRET")
}

func GetPort() int {
	return viper.GetInt("PORT")
}

func GetEndPoint() string {
	return viper.GetString("ENDPOINT_NODE_API")
}
