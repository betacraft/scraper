package scraper

import (
	"testing"
)

func TestPlayStoreScraper(t *testing.T) {
	app, err := Scrape("https://play.google.com/store/apps/details?id=com.simpl.android")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(app)
}
