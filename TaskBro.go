package main

import (
	// "github.com/gin-gonic/gin"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("app") // no need to include file extension
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		fmt.Println("DBName", viper.Get("dbname"))
	}

}
