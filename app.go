package scraper

type App struct {
	Name        string
	IconUrl     string
	Price       string
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
	RatingValue     float64
	RatingCount     int
	ContentRating   string

	Downloads string

	ApkDownloadUrl  string
	ApkDownloadSize string
	ApkDownloadMD5  string
}
