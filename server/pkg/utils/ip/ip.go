package ip

import (
	"net"

	"github.com/kuops/go-example-app/server/pkg/log"
)

func GetIPAddress() string {
	var ip string
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Error("get ip address failed")
		return "127.0.0.1"
	}

	defer conn.Close()

	if conn != nil {
		ip = conn.LocalAddr().(*net.UDPAddr).IP.String()
	}

	return ip
}
