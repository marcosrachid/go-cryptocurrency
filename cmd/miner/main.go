package main

import (
	"go-cryptocurrency/internal/handler"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = handler.WalletStart()
	if err != nil {
		panic(err)
	}
}
