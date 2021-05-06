package main

import (
	"context"
	"errors"
	"github.com/google/safebrowsing"
	log "github.com/sirupsen/logrus"
	"os"
)

type sbAPI struct {
	*safebrowsing.SafeBrowser
}

var (
	ErrSafeBrowsingAPIKeyNotFound = errors.New("no API_KEY provided in env vars")
)

const SAFEBROWSINGAPIKEY = "GOOGLE_SAFEBROWSING_API_KEY"

func createSafeBrowsingAPI() (*sbAPI, error) {
	apiKey, found := os.LookupEnv(SAFEBROWSINGAPIKEY)
	if !found {
		log.Println("SafeBrowsing API key is not found")
		return nil, ErrSafeBrowsingAPIKeyNotFound
	}

	config := safebrowsing.Config{
		APIKey: apiKey,
		Logger: os.Stdout,
	}
	sb, err := safebrowsing.NewSafeBrowser(config)
	log.Println("Try to create SafeBrowsing API")
	return &sbAPI{sb}, err
}

func (sb *sbAPI) CheckUrls(ctx context.Context, urlRecords []URLRecord) (bool, error) {
	urls := make([]string, 0)
	for _, urlRecord := range urlRecords {
		urls = append(urls, urlRecord.Url)
	}
	log.Println("Check URLs by SafeBrowsing API")
	threats, err := sb.LookupURLsContext(ctx, urls)

	if len(threats[0]) != 0 {
		log.Println("Threats were found")
		return true, err
	}
	log.Println("Threats were not found")
	return false, err
}
