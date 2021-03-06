package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	cfg := Config{}

	flag.StringVar(&cfg.Path, "path", "", "Path of the documentation")

	flag.BoolVar(&cfg.Debug, "debug", false, "Debug mode")

	version := flag.Bool("v", false, "Show version.")
	help := flag.Bool("h", false, "Show this help.")

	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
	}

	if *version {
		displayVersion()
		return
	}

	if flag.NArg() > 0 {
		usage()
	}

	err := validate(cfg)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}

	err = run(cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func usage() {
	_, _ = os.Stderr.WriteString(fmt.Sprintf("Seo (%s)\n\nFlags:\n", version))
	flag.PrintDefaults()
	os.Exit(2)
}

func validate(cfg Config) error {
	if strings.TrimSpace(cfg.Path) == "" {
		return errors.New("path is required")
	}

	return nil
}
