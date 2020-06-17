package main

import (
	"os"

	"go-cryptocurrency/internal/handler"
	"go-cryptocurrency/internal/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

var Blockchain []models.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = handler.KeyStart()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Logger.Fatal(e.Start(":" + os.Getenv("REST_PORT")))
}
