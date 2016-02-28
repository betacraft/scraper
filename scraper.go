package scraper

import (
	"errors"
	"log"
	"net/url"

	"github.com/betacraft/goquery"
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

func ScrapeDoc(rawUrl string, doc *goquery.Document) (*App, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	log.Println("Parsing", uri.Host)
	switch uri.Host {
	case "play.google.com":
		return parsePlayStorePage(doc)
	case "www.androiddrawer.com":
		fallthrough
	case "androiddrawer.com":
		return parseAndroidDrawerPage(doc)
	}
	return nil, errors.New("Site is not supported by parser")
}
