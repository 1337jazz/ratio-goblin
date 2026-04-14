package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/1337jazz/ratio-goblin/internal/config"
	"github.com/1337jazz/ratio-goblin/internal/scraper"
	"github.com/1337jazz/ratio-goblin/internal/updater"
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
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("Failed to load config - did you run `ratiogoblin init` first?\n\nERROR:", err)
			os.Exit(1)
		}
		updateCh := make(chan bool, 1)
		go func() { updateCh <- updater.HasUpdate() }()

		s := scraper.NewScraper(cfg)
		ratio := s.ScrapeRatio()

		// Block and wait for the update check to finish before printing the ratio
		if <-updateCh {
			fmt.Println(ratio + " *")
		} else {
			fmt.Println(ratio)
		}
	case "update":
		if err := updater.Update(); err != nil {
			if errors.Is(err, updater.UpdateCancelled) {
				fmt.Println("Update cancelled")
				return
			}
			if errors.Is(err, updater.AlreadyUpToDate) {
				fmt.Println("Already up to date")
				return
			}
			fmt.Println("Failed to update:", err)
			os.Exit(1)
		}
		fmt.Println("Updated successfully - restart the ratiogoblin process in your bar")
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
  update          Update to the latest release
  version         Display version information`)
}

func printVersion() {
	fmt.Printf("ratiogoblin version %s\n", version.Version)
}
