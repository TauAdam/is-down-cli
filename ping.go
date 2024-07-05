package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func Ping(domain, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("DOWN  %v %v", err, address)
	} else {
		status = fmt.Sprintf("UP   %v %v", conn.RemoteAddr(), conn.LocalAddr())
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Close connection")
		}
	}(conn)
	return status
}
