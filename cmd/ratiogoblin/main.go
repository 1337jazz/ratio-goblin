package main

import (
	"fmt"
	"os"

	"github.com/1337jazz/ratio-goblin/internal/config"
	"github.com/1337jazz/ratio-goblin/internal/scraper"
)

func main() {

	if len(os.Args) < 2 {
		// TODO: print help
		fmt.Println("Usage: ratiogoblin <command>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		err := config.InitConfig()
		if err != nil {
			fmt.Println("Failed to create config:", err)
			os.Exit(1)
		}
	case "run":
		s := scraper.NewScraper(config.LoadConfig())
		ratio := s.ScrapeRatio()
		fmt.Println(ratio)
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}

}
