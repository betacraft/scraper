package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"strings"
)

func parseAndroidDrawer(url *url.URL) (*App, error) {
	app := new(App)
	doc, err := goquery.NewDocument(url.String())
	if err != nil {
		return nil, err
	}
	var playurl string
	doc.Find(".playurl").Each(func(i int, s *goquery.Selection) {
		playurl, _ = s.Attr("href")
		log.Println("playurl", playurl)
	})
	app, err = Scrape(playurl)
	if err != nil {
		return nil, err
	}
	doc.Find(".changelog-wrap").Children().Each(func(i int, s *goquery.Selection) {
		attr, _ := s.Attr("class")
		if attr != "download-wrap" {
			return
		}
		href, _ := s.Children().Attr("href")
		log.Println(href)
		app.ApkDownloadUrl = href
		// getting other info
		s.Children().Children().Each(func(i int, s *goquery.Selection) {
			attr, _ := s.Attr("class")
			switch attr {
			case "download-md5":
				app.ApkDownloadMD5 = strings.Replace(s.Text(), "MD5: ", "", -1)
			case "download-size":
				app.ApkDownloadSize = s.Text()
			}
		})
	})
	return app, nil
}
