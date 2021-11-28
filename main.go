package main

import (
	"fmt"

	Config "github.com/ahmedsharyo/AlertService/config"
	Models "github.com/ahmedsharyo/AlertService/models"
	"github.com/gin-gonic/gin"

	//"github.com/ahmedsharyo/AlertService/Routes"

	"github.com/jinzhu/gorm"

	"github.com/ahmedsharyo/AlertService/grpc_service"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Alert{})
	server := gin.New()

	//Routes.Setup(app)
	go grpc_service.Run()

	port := "8080"
	server.Run(":" + port)

}
