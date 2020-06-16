package main

import (
	"log"
	"os"
	"time"

	"go-cryptocurrency/internal/models"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

var Blockchain []models.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := models.GenesisBlock{0, t.Unix(), os.Getenv("COINBASE"), "", "", 0}
		genesisBlock.Hash = genesisBlock.CalculateHash()
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()

	log.Println(Blockchain)

	e := echo.New()
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
