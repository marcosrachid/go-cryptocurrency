package network

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const (
	END = "\r\n\r\n"
)

func SocketServer(port string) {

	listen, err := net.Listen("tcp4", ":"+port)

	if err != nil {
		log.Fatalf("Socket listen port %s failed,%s", port, err)
		os.Exit(1)
	}

	defer listen.Close()

	log.Printf("Begin listen port: %s", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go func(conn net.Conn) {

			defer conn.Close()

			var (
				buf = make([]byte, 1024)
				r   = bufio.NewReader(conn)
				w   = bufio.NewWriter(conn)
			)

		ILOOP:
			for {
				n, err := r.Read(buf)
				data := string(buf[:n])

				switch err {
				case io.EOF:
					break ILOOP
				case nil:
					Handle(data, w)
					if isTransportOver(data) {
						break ILOOP
					}

				default:
					log.Fatalf("Receive data failed:%s", err)
					return
				}

			}
		}(conn)
	}

}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, END)
	return
}
