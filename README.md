# Twitch Emote Scraper
This is a tool to download a lot of twitch emotes and their names. This is done by generating emote set ids starting from 1 and going up, and downloading the associated emotes from twitch.

## Twitch Credentials
To use the Twitch API you will need a Client ID and Client Secret. This can be achieved by creating an application at [Twitch Developer Portal](https://dev.twitch.tv/)

## Installation
To build this tool you will need to install the [Go Programming Language](https://go.dev/dl/).

Then you will need to clone this repository and run `go run ./cmd/main.go`.

You can also directly install the CLI tool using:

`go install github.com/sha65536/twitchemotes/cmd/emotescraper@latest`.

## Usage
I've included a small CLI tool in the cmd folder:
```
NAME:
   EmoteScraper - Scrape twitch emotes

USAGE:
   EmoteScraper [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --client-id value, --ci value      client-id to scrape with (default: from env)
   --client-secret value, --cs value  secret for the client (default: from env)
   --size value                       size of the emotes (1.0, 2.0, 3.0) (default: 3.0)
   --logging value, -l value          logging to file (default: true)
   --start value                      emote set id to start with (default: 0)
   --output value, -o value           output folder (default: output/)
   --help, -h                         show help
```
You can either pass variables using the CLI or using environment variables / .env file:
```
CLIENT_ID="yourtwitchid"
CLIENT_SECRET="yourtwitchsecret"
SIZE="3.0"
LOGGING="true"
OUTPUT="output"
START="0"
```