package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func Ping(domain, port, protocol string, timeout time.Duration, verbose bool) {
	address := domain + ":" + port
	conn, err := net.DialTimeout(protocol, address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("DOWN  %v %v", err, address)
	} else {
		status = fmt.Sprintf("UP   %v %v", conn.RemoteAddr(), conn.LocalAddr())
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				log.Println("Close connection")
			}
		}(conn)
	}
	if verbose {
		log.Println("Verbose mode enabled")
		log.Println("Domain:", domain)
		log.Println("Port:", port)
		log.Println("Protocol:", protocol)
		log.Println("Timeout:", timeout)
		log.Println("Status:", status)
	} else {
		log.Println(status)
	}
}
