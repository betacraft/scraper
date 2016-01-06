# Scraper for parsing app information from Google Play Store

## How to use?
```go
import "github.com/betacraft/scraper"

func xyz(){
   app, err := scraper.Scrape("https://play.google.com/store/apps/details?id=com.simpl.android")
   if err != nil {
       //parsing failed
   }
}
```
## App struct
```go
type App struct {
	Name        string
	IconUrl     string
	Price       int
	Description string
	LastUpdated string
	FileSize    string

	OperatingSystem string

	VersionName string

	ScreenshotUrls []string

	Author    string
	AuthorUrl string

	Genre string

	AggregateRating string
	RatingValue     int
	RatingCount     int
	ContentRating   string

	Downloads string
}
```

## Supported sites
1. Google Play Store
