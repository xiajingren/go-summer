package conf

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Api struct {
		Gin_Mode string
		Gin_Port string
		Jwt_Key  []byte
		Jwt_Exp  int64
	}
	MySql struct {
		Host,
		Post,
		Db_Name,
		User,
		Password string
	}
}

var Conf Config

func InitConfig() {
	viper.SetConfigName("conf")       // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../../conf") // optionally look for config in the working directory
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %w \n", err)
	}

	setConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		setConfig()
	})
	viper.WatchConfig()
}

func setConfig() {
	Conf.Api.Gin_Mode = viper.GetString("api.gin_mode")
	Conf.Api.Gin_Port = viper.GetString("api.gin_port")
	Conf.Api.Jwt_Key = []byte(viper.GetString("api.jwt_key"))
	Conf.Api.Jwt_Exp = viper.GetInt64("api.jwt_exp")

	Conf.MySql.Host = viper.GetString("mysql.host")
	Conf.MySql.Post = viper.GetString("mysql.post")
	Conf.MySql.Db_Name = viper.GetString("mysql.db_name")
	Conf.MySql.User = viper.GetString("mysql.user")
	Conf.MySql.Password = viper.GetString("mysql.password")
}
