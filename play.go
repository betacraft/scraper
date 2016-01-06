package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func parsePlayStore(url *url.URL) (*App, error) {
	log.Println("Parsing", url.String())
	app := new(App)
	doc, err := goquery.NewDocument(url.String())
	if err != nil {
		return nil, err
	}
	log.Println("Getting all itemprops")
	itemprops := make([]goquery.Selection, 0)
	// adding all itemprops
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		getItemprop(s, &itemprops)
	})
	// for getting logo url
	doc.Find(".cover-image").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("src")
		if !ok {
			return
		}
		app.IconUrl = url
	})
	// sanitizing data
	for _, prop := range itemprops {
		itemprop, _ := prop.Attr("itemprop")
		switch itemprop {
		case "name":
			// checking parent to differ between author name and app name
			parent, _ := prop.Html()
			if strings.Contains(parent, "id-app-title") {
				log.Println("name", prop.Text())
				app.Name = prop.Text()
				continue
			}
			if prop.Text() == "" {
				continue
			}
			log.Println("author", prop.Text())
			app.Author = prop.Text()
		case "genre":
			log.Println("genre", prop.Text())
			app.Genre = prop.Text()
		case "price":
			price, _ := prop.Attr("content")
			log.Println("price", price)
			app.Price, err = strconv.Atoi(price)
			if err != nil {
				return nil, err
			}
		case "screenshot":
			url, _ := prop.Attr("src")
			log.Println("screenshot", url)
			app.ScreenshotUrls = append(app.ScreenshotUrls, url)
		case "description":
			log.Println("description", prop.Text())
			app.Description = prop.Text()
		case "aggregateRating":
			log.Println("aggregateRating", prop.Text())
			app.AggregateRating = prop.Text()
		case "ratingValue":
			value, _ := prop.Attr("content")
			log.Println("ratingValue", value)
			app.RatingValue, err = strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
		case "ratingCount":
			count, _ := prop.Attr("content")
			log.Println("ratingCount", count)
			app.RatingCount, err = strconv.Atoi(count)
			if err != nil {
				return nil, err
			}
		case "datePublished":
			log.Println("datePublished", prop.Text())
			app.LastUpdated = prop.Text()
		case "fileSize":
			log.Println("fileSize", prop.Text())
			app.FileSize = prop.Text()
		case "numDownloads":
			log.Println("numDownloads", prop.Text())
			app.Downloads = prop.Text()
		case "softwareVersion":
			log.Println("softwareVersion", prop.Text())
			app.VersionName = prop.Text()
		case "operatingSystems":
			log.Println("operatingSystems", prop.Text())
			app.OperatingSystem = prop.Text()
		case "contentRating":
			log.Println("contentRating", prop.Text())
			app.ContentRating = prop.Text()
		}
	}
	return app, nil
}

func getItemprop(s *goquery.Selection, array *[]goquery.Selection) {
	itemprop, ok := s.Attr("itemprop")
	if !ok {
		return
	}
	log.Println("Adding", itemprop)
	*array = append(*array, *s)
}
