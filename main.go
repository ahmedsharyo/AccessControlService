package main

import (
	"fmt"

	Config "github.com/ahmedsharyo/AccessControlService/config"
	Models "github.com/ahmedsharyo/AccessControlService/models"
	"github.com/gin-gonic/gin"

	//"github.com/ahmedsharyo/AlertService/Routes"

	"github.com/jinzhu/gorm"

	"github.com/ahmedsharyo/AccessControlService/grpc_service"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.SecurityMangerCameras{})
	server := gin.New()

	//Routes.Setup(app)
	go grpc_service.Run()

	port := "8081"
	server.Run(":" + port)

}
