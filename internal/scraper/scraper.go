package scraper

import (
	"fmt"
	"net/http"

	"github.com/1337jazz/ratio-goblin/internal/config"
	"github.com/PuerkitoBio/goquery"
)

const URL = "https://iptorrents.com/user.php"

type Scraper interface {

	// ScrapeRatio scrapes the user's ratio from the website.
	ScrapeRatio() string
}

type scraper struct {
	uid  string
	pass string
}

// NewScraper creates a new Scraper instance with the given configuration.
func NewScraper(config *config.Config) Scraper {
	return &scraper{
		uid:  config.CookieUID,
		pass: config.CookiePass,
	}
}

func (s *scraper) ScrapeRatio() string {

	// Create HTTP request
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	// Add the "u" query parameter (the uid)
	q := req.URL.Query()
	q.Add("u", s.uid)
	req.URL.RawQuery = q.Encode()

	// Set cookies (uid and pass)
	req.Header.Set("Cookie", fmt.Sprintf("uid=%s; pass=%s", s.uid, s.pass))

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if resp.StatusCode != http.StatusOK {
			return fmt.Sprintf("ERROR: %s", err.Error())
		}
		return fmt.Sprintf("ERROR: %s", err.Error())
	}
	defer resp.Body.Close()

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	// Extract the ratio value using the CSS selector
	value := doc.Find(".al > font:nth-child(1) > font:nth-child(1)").Text()

	return value
}
