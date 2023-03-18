package twitchemotes

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nicklaw5/helix/v2"
)

const (
	SizeSmall  = "1.0"
	SizeMedium = "2.0"
	SizeLarge  = "3.0"
)

const urlTemplate = "https://static-cdn.jtvnw.net/emoticons/v2/%s/static/light/%s"

type EmoteScraper struct {
	Current []byte
	Client  *helix.Client
	Size    string
	Logging bool
}

func MakeScraper(clientId, clientSecret, size string, logging bool) (*EmoteScraper, error) {
	if logging {
		f, err := os.OpenFile("scrape.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		log.SetOutput(io.MultiWriter(f, os.Stdout))
	}
	client, err := helix.NewClient(&helix.Options{
		ClientID:     clientId,
		ClientSecret: clientSecret,
	})
	if err != nil {
		return nil, err
	}
	resp, err := client.RequestAppAccessToken([]string{"user:read:email"})
	if err != nil {
		return nil, err
	}
	client.SetAppAccessToken(resp.Data.AccessToken)

	return &EmoteScraper{Current: []byte{'0'}, Client: client, Size: size, Logging: logging}, nil
}

func (es *EmoteScraper) Log(v ...any) {
	if es.Logging {
		log.Println(v...)
	}
}

func (es *EmoteScraper) SetCurrent(start string) {
	es.Current = []byte(start)
	es.Log("set start to", start)
}

func (es *EmoteScraper) StartScraping(output string) error {
	var iter uint64
	if err := os.Mkdir(output, 0666); err != nil && !os.IsExist(err) {
		es.Log("error opening output", err)
		return err
	}
	es.Log("started scraping")
	for {
		sets := es.Generate()
		es.Log("trying", sets[0], "to", sets[len(sets)-1])
		batch, err := es.Client.GetEmoteSets(&helix.GetEmoteSetsParams{
			EmoteSetIDs: sets,
		})
		if err != nil {
			es.Log("error getting emote set", err)
			return err
		}
		for _, e := range batch.Data.Emotes {
			iter++
			es.Log(iter, e.OwnerID, e.Name)
			if err := DownloadEmote(output, e.Name, e.ID, es.Size); err != nil {
				es.Log("error downloading emote", e.ID, err)
			}
		}
	}
}

func DownloadEmote(path, name, id, size string) error {
	var filename = filepath.Join(path, name+".png")
	var url = fmt.Sprintf(urlTemplate, id, size)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	return err
}
