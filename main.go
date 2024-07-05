package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Name:  "Check whether resource down or not",
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
				Usage:       "set port number",
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
				Usage: "set protocol (tcp, udp)",
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
			&cli.StringFlag{
				Name:        "log",
				Usage:       "set log file",
				Value:       "log.log",
				Aliases:     []string{"l"},
				DefaultText: "log.log",
			},
		},
		Action: func(c *cli.Context) error {
			logFile := c.String("log")
			f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				return err
			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {

				}
			}(f)
			log.SetOutput(f)

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
				Ping(domain, port, protocol, timeout, verbose)
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
