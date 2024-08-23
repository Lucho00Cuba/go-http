package main

import (
	"fmt"
	"os"

	"github.com/lucho00cuba/go-http/pkg/server"
	"github.com/lucho00cuba/go-http/version"
	"github.com/urfave/cli/v2"
)

type Config struct {
	ServerPort string
	ServerHost string
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func main() {
	app := &cli.App{
		Name:    "go-server",
		Usage:   "A CLI tool for running a Go server or client",
		Version: version.VERSION,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "server-port",
				Usage:   "Port for the server to listen on",
				EnvVars: []string{"SERVER_PORT"},
				Value:   GetEnv("SERVER_PORT", "3000"),
			},
			&cli.StringFlag{
				Name:    "server-host",
				Usage:   "Host for the server",
				EnvVars: []string{"SERVER_HOST"},
				Value:   GetEnv("SERVER_HOST", "localhost"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Run the server",
				Action: func(c *cli.Context) error {
					config := &Config{
						ServerPort: c.String("server-port"),
						ServerHost: c.String("server-host"),
					}
					fmt.Printf("Running server with config: %+v\n", config)
					srv := server.NewServer(config.ServerPort)
					srv.Run()
					return nil
				},
			},
			{
				Name:  "client",
				Usage: "Run the client",
				Action: func(c *cli.Context) error {
					// TO-DO
					fmt.Println("Running client...")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
