package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sha65536/twitchemotes"
	"github.com/urfave/cli/v2"
)

var (
	clientId     string                          // CLIENT_ID
	clientSecret string                          // CLIENT_SECRET
	size         string = twitchemotes.SizeLarge // SIZE
	logging      string = "true"                 // LOGGING
	start        string = "0"                    // START
	output       string = "output"               // OUTPUT
)

func main() {
	getDefaults()
	app := &cli.App{
		Name:  "EmoteScraper",
		Usage: "Scrape twitch emotes",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "client-id",
				Aliases:     []string{"ci"},
				Usage:       "client-id to scrape with",
				DefaultText: "from env",
				Value:       clientId,
				Destination: &clientId,
			},
			&cli.StringFlag{
				Name:        "client-secret",
				Aliases:     []string{"cs"},
				Usage:       "secret for the client",
				DefaultText: "from env",
				Value:       clientSecret,
				Destination: &clientSecret,
			},
			&cli.StringFlag{
				Name:        "size",
				Usage:       "size of the emotes (1.0, 2.0, 3.0)",
				Value:       size,
				DefaultText: "3.0",
				Destination: &size,
			},
			&cli.StringFlag{
				Name:        "logging",
				Aliases:     []string{"l"},
				Usage:       "logging to file",
				Value:       logging,
				DefaultText: "true",
				Destination: &logging,
			},
			&cli.StringFlag{
				Name:        "start",
				Usage:       "emote set id to start with",
				Value:       start,
				DefaultText: "0",
				Destination: &start,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "output folder",
				Value:       output,
				DefaultText: "output/",
				Destination: &output,
			},
		},
		Action: func(*cli.Context) error {
			scraper, err := twitchemotes.MakeScraper(
				clientId,
				clientSecret,
				size,
				strings.ToLower(logging) == "true",
			)
			if err != nil {
				log.Fatal(err)
			}
			scraper.SetCurrent(start)
			if err := scraper.StartScraping(output); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getDefaults() {
	godotenv.Load()
	if val := os.Getenv("CLIENT_ID"); val != "" {
		clientId = val
	}
	if val := os.Getenv("CLIENT_SECRET"); val != "" {
		clientSecret = val
	}
	if val := os.Getenv("SIZE"); val != "" {
		size = val
	}
	if val := os.Getenv("LOGGING"); val != "" {
		logging = val
	}
	if val := os.Getenv("START"); val != "" {
		start = val
	}
	if val := os.Getenv("OUTPUT"); val != "" {
		output = val
	}
}
