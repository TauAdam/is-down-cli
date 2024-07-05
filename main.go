package is_down_cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "is_down_cli",
		Usage: "download cli tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Usage:    "domain name",
				Aliases:  []string{"d"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Usage:    "port number",
				Aliases:  []string{"p"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if port == "" {
				port = "8080"
			}
			status := Ping(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
