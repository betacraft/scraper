package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
)

func parsePlayStorePage(doc *goquery.Document) (*App, error) {
	var err error
	app := new(App)
	itemprops := make([]goquery.Selection, 0)
	// adding all itemprops
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		getItemprop(s, &itemprops)
	})
	// sanitizing data
	for _, prop := range itemprops {
		itemprop, _ := prop.Attr("itemprop")
		switch itemprop {
		case "image":
			app.IconUrl, _ = prop.Attr("src")
			app.IconUrl = strings.Split(app.IconUrl, "=w")[0]
			app.IconUrl += "=w512"
		case "name":
			// checking parent to differ between author name and app name
			parent, _ := prop.Html()
			if strings.Contains(parent, "id-app-title") {
				app.Name = prop.Text()
				continue
			}
			if prop.Text() == "" {
				continue
			}
			app.Author = prop.Text()
		case "genre":
			app.Genre = prop.Text()
		case "price":
			price, _ := prop.Attr("content")
			app.Price = price
		case "screenshot":
			url, _ := prop.Attr("src")
			app.ScreenshotUrls = append(app.ScreenshotUrls, url)
		case "description":
			app.Description = prop.Text()
		case "aggregateRating":
			app.AggregateRating = prop.Text()
		case "ratingValue":
			value, _ := prop.Attr("content")
			app.RatingValue, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
		case "ratingCount":
			count, _ := prop.Attr("content")
			app.RatingCount, err = strconv.Atoi(count)
			if err != nil {
				return nil, err
			}
		case "datePublished":
			app.LastUpdated = prop.Text()
		case "fileSize":
			app.FileSize = prop.Text()
		case "numDownloads":
			app.Downloads = prop.Text()
		case "softwareVersion":
			app.VersionName = prop.Text()
		case "operatingSystems":
			app.OperatingSystem = prop.Text()
		case "contentRating":
			app.ContentRating = prop.Text()
		}
	}
	return app, nil
}

func parsePlayStore(url *url.URL) (*App, error) {
	doc, err := goquery.NewDocument(url.String())
	if err != nil {
		return nil, err
	}
	return parsePlayStorePage(doc)
}

func getItemprop(s *goquery.Selection, array *[]goquery.Selection) {
	_, ok := s.Attr("itemprop")
	if !ok {
		return
	}
	*array = append(*array, *s)
}
