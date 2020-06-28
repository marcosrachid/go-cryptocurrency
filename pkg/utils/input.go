package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	var s string
	if s, _ = reader.ReadString('\n'); true {
		s = strings.TrimSpace(s)
	}
	return s
}
