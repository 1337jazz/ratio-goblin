package scraper

import (
	"fmt"
	"net/http"

	"github.com/1337jazz/ratio-goblin/internal/config"
	"github.com/PuerkitoBio/goquery"
)

const url = "https://iptorrents.com/user.php"

type Scraper interface {
	ScrapeRatio() string
}

type scraper struct {
	uid  string
	pass string
}

func NewScraper(config *config.Config) Scraper {
	return &scraper{
		uid:  config.CookieUID,
		pass: config.CookiePass,
	}
}

func (s *scraper) ScrapeRatio() string {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	q := req.URL.Query()
	q.Add("u", s.uid)

	req.URL.RawQuery = q.Encode()
	req.Header.Set("Cookie", "uid="+s.uid+"; pass="+s.pass)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if resp.StatusCode != http.StatusOK {
			return fmt.Sprintf("ERROR: %s", err.Error())
		}
		return fmt.Sprintf("ERROR: %s", err.Error())
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	value := doc.Find(".al > font:nth-child(1) > font:nth-child(1)").Text()

	return value
}
