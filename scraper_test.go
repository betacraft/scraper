package scraper

import (
	"log"
	"testing"
)

func TestPlayStoreScraper(t *testing.T) {
	app, err := Scrape("https://play.google.com/store/apps/details?id=com.linkedin.android")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	log.Println(app)
}

func TestAndroidDrawerScraper(t *testing.T) {
	app, err := Scrape("http://www.androiddrawer.com/19821/download-aldiko-book-reader-app-apk/")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	log.Println(app)
}
