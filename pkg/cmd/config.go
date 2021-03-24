package cmd

import (
	"fmt"
	"sort"

	"github.com/urfave/cli/v2"
)

type Config struct {
	Port                      int
	BindAddress               string
	WriteTimeout, ReadTimeout int
	MappingDirectory          string
}

func NewApp() (*Config, *cli.App) {
	config := &Config{}
	app := cli.NewApp()
	app.Name = "mimic"
	app.Usage = "Your friendly and simple HTTP mock server"
	app.Version = Version

	app.Action = func(c *cli.Context) error {
		fmt.Println("Starting mimic")
		return nil
	}

	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:        "port",
			Value:       8080,
			Usage:       "Set the HTTP/S port number",
			Destination: &config.Port,
		},
		&cli.StringFlag{
			Name:        "bind-address",
			Value:       "localhost",
			Usage:       "The IP address the mimic server should serve from.",
			Destination: &config.BindAddress,
		},
		&cli.IntFlag{
			Name:        "read-timeout",
			Value:       15,
			Usage:       "Maximum duration for reading the entire request, including the body  (in seconds)",
			Destination: &config.ReadTimeout,
		},
		&cli.IntFlag{
			Name:        "write-timeout",
			Value:       15,
			Usage:       "Maximum duration before timing out writes of the response (in seconds)",
			Destination: &config.WriteTimeout,
		},
		&cli.StringFlag{
			Name:        "mapping",
			Aliases:     []string{"m"},
			Usage:       "Load mapping file from directory",
			Destination: &config.MappingDirectory,
			Required:    true,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	return config, app
}
