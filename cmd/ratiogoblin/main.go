package main

import (
	"fmt"
	"os"

	"github.com/1337jazz/ratio-goblin/internal/config"
	"github.com/1337jazz/ratio-goblin/internal/scraper"
	"github.com/1337jazz/ratio-goblin/internal/version"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		printUsage()
	case "init":
		err := config.InitConfig()
		if err != nil {
			fmt.Println("Failed to create config:", err)
			os.Exit(1)
		}
		fmt.Println("Config file created successfully.")
	case "run":
		s := scraper.NewScraper(config.LoadConfig())
		ratio := s.ScrapeRatio()
		fmt.Println(ratio)
	case "version":
		printVersion()
	default:
		fmt.Println("Unknown command:", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Println(`Usage: ratiogoblin <command>
Commands:
  help            Show this help message
  init            Initialize configuration file with default settings
  run             Execute the ratio scraper and display the output
  version         Display version information`)
}

func printVersion() {
	fmt.Printf("ratiogoblin version %s\n", version.Version)
}
