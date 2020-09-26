package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
)

func main() {
	port := 24224
	host := "127.0.0.1"
	requests := 1
	concurrency := 1

	cli.HelpFlag = &cli.BoolFlag{
		Name: "help",
	}

	app := &cli.App{
		Name:        "bulk-fluent-cat",
		Description: "send records",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   port,
				Usage:   "fluent tcp port (default: " + strconv.Itoa(port) + ")",
			},
			&cli.StringFlag{
				Name:    "host",
				Aliases: []string{"h"},
				Value:   host,
				Usage:   "fluent host (default: " + host + ")",
			},
			&cli.IntFlag{
				Name:    "requests",
				Aliases: []string{"n"},
				Value:   requests,
				Usage:   "Number of requests to send messages (default: " + strconv.Itoa(requests) + ")",
			},
			&cli.IntFlag{
				Name:    "concurrency",
				Aliases: []string{"c"},
				Value:   concurrency,
				Usage:   "Number of multiple requests to perform at a time (default: " + strconv.Itoa(concurrency) + ")",
			},
		},
		Action: func(context *cli.Context) error {
			tag := context.Args().Get(0)
			if tag == "" {
				return fmt.Errorf("specify tag")
			}
			fluentCat(context.String("host"), context.Int("port"), context.Int("requests"), context.Int("concurrency"), tag, "message")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
