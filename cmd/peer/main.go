package main

import (
	"log"
	"os"

	"go-cryptocurrency/internal/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

var Blockchain []models.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Logger.Fatal(e.Start(":" + os.Getenv("REST_PORT")))
}
