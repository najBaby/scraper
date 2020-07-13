package scraper

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
)

// Scraper is
type Scraper struct {
	//remote config
	remote *Remote
}

//cheick if the request has redirected
func checkRedirect(req *http.Request, via []*http.Request) error {
	return nil
}

// NewScraper Init new scrapeur with cookies and user-agent
func NewScraper() (*Scraper, error) {
	s := new(Scraper)
	jar, err := cookiejar.New(nil) // &cookiejar.Options{PublicSuffixList: publicsuffix.List} "golang.org/x/net/publicsuffix"
	if err != nil {
		return nil, err
	}
	s.remote = NewRemote(Config{
		Jar:           jar,
		CheckRedirect: checkRedirect,
	})
	return s, nil
}

// Scraping Shedule scraping function
func (scraper *Scraper) Scraping(url string) (*goquery.Document, error) {
	res, err := scraper.remote.GET(Options{
		URL: url,
		Header: map[string]string{
			"Connection":      "keep-alive",
			"Accept-Language": "fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5",
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
			"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
		},
	})

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// goquery parses response body
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
