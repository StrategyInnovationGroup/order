package main

import (
	"fmt"
	"order/pkg/handler"
	"order/pkg/initializer"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
)

var dbConn *gorm.DB

func init() {
	fmt.Println("Loading Env vars ...")
	initializer.LoadEnvVars()
	fmt.Println("Env vars loaded successfully , initialization to DB started ...")
	dbConn = initializer.DBConnection()
	fmt.Println("DB connection completed. Migration Running ...")
	error := initializer.RunDBMigration()

	if error == migrate.ErrNoChange {
		fmt.Println("No change detected in migration ... ")
		return
	}

	if error != nil {
		fmt.Println("DB migration failed ... ")
		return
	}

	fmt.Println("Application started successfully ... ")
}

func main() {
	router := gin.Default()

	routerGroup := router.Group("/api/v1/")
	routerGroup.GET("ping", handlePing)
	orderRouterGroup := routerGroup.Group("/order")
	handler.Controllers(orderRouterGroup)
	router.Run()
}

func handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
