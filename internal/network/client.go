package network

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/marcosrachid/go-cryptocurrency/internal/global"
	"github.com/marcosrachid/go-cryptocurrency/internal/models"
)

func SocketClient(ip, port string, request models.CLIRequest) {
	addr := strings.Join([]string{ip, port}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	data, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn.Write([]byte(string(data) + global.END))

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	fmt.Printf("%s", buff[:n])

}
