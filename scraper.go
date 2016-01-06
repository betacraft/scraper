package scraper

import (
	"errors"
	"log"
	"net/url"
)

func Scrape(rawUrl string) (*App, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	log.Println(uri.Host)
	switch uri.Host {
	case "play.google.com":
		return parsePlayStore(uri)
	case "www.androiddrawer.com":
		fallthrough
	case "androiddrawer.com":
		return parseAndroidDrawer(uri)
	}
	return nil, errors.New("Site is not supported by parser")
}
