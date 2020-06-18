package main

import (
	"os"

	"go-cryptocurrency/internal/handler"
	"go-cryptocurrency/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Blockchain []models.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = handler.WalletStart()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Run(":" + os.Getenv("REST_PORT"))
}
