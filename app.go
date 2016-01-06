package scraper

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
