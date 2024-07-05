package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Config struct {
	Domain   string
	Port     string
	Protocol string
	Timeout  time.Duration
	Verbose  bool
}

func Ping(config Config) {
	address := fmt.Sprintf("%s:%s", config.Domain, config.Port)
	conn, err := net.DialTimeout(config.Protocol, address, config.Timeout)
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
	if config.Verbose {
		log.Println("Verbose mode enabled")
		log.Printf("Domain: %s\n", config.Domain)
		log.Printf("Port: %s\n", config.Port)
		log.Printf("Protocol: %s\n", config.Protocol)
		log.Printf("Timeout: %v\n", config.Timeout)
		log.Printf("Status: %s\n", status)
	} else {
		log.Println(status)
	}
}
