package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Name:  "is_down_cli",
		Usage: "download cli tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Usage:    "Specify domain name",
				Aliases:  []string{"d"},
				Required: true,
			},
			&cli.StringFlag{
				Name:        "port",
				Usage:       "port number",
				Aliases:     []string{"p"},
				DefaultText: "80",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Usage:   "enable verbose mode",
				Aliases: []string{"v"},
			},
			&cli.DurationFlag{
				Name:  "timeout",
				Usage: "set timeout value",
				Value: 5 * time.Second,
			},
			&cli.StringFlag{
				Name:  "protocol",
				Usage: "set protocol (tcp, udp, icmp)",
				Value: "tcp",
			},
			&cli.BoolFlag{
				Name:    "repeat",
				Usage:   "enable repeat mode",
				Aliases: []string{"r"},
			},
			&cli.DurationFlag{
				Name:  "interval",
				Usage: "set interval for repeat mode",
				Value: 1 * time.Minute,
			},
		},
		Action: func(c *cli.Context) error {
			domain := c.String("domain")
			port := c.String("port")
			if len(port) == 0 {
				port = "80"
			}
			verbose := c.Bool("verbose")
			timeout := c.Duration("timeout")
			protocol := c.String("protocol")
			repeat := c.Bool("repeat")
			interval := c.Duration("interval")

			for {
				status := Ping(domain, port, protocol, timeout, verbose)
				fmt.Println(status)
				if !repeat {
					break
				}
				time.Sleep(interval)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
