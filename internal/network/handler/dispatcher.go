package handler

import (
	"bufio"
	"log"
)

func DispatcherHandler(data string, writer *bufio.Writer) {
	log.Println(data)
}
