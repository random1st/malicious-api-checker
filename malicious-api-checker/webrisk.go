package main

import (
	"context"
	"errors"
	"github.com/google/webrisk"
	log "github.com/sirupsen/logrus"
	"os"
)

const WEBRIKSKAPIKEY = "GOOGLE_WEBRISK_API_KEY"

type wrAPI struct {
	*webrisk.WebriskClient
}

var (
	ErrWebRiskAPIKeyNotFound = errors.New("no API_KEY provided in env vars")
)

func createWebRiskAPI() (*wrAPI, error) {
	apiKey, found := os.LookupEnv(WEBRIKSKAPIKEY)
	if !found {
		log.Println("WebRisk API key is not found")
		return nil, ErrWebRiskAPIKeyNotFound
	}

	config := webrisk.Config{
		APIKey: apiKey,
		Logger: os.Stdout,
	}

	wr, err := webrisk.NewWebriskClient(config)
	log.Println("Try to create WebRisk API ")
	return &wrAPI{wr}, err
}
func (wr *wrAPI) CheckUrls(ctx context.Context, urlRecords []URLRecord) ([]URLThreat, error) {
	urls := make([]string, 0)
	for _, urlRecord := range urlRecords {
		urls = append(urls, urlRecord.Url)
	}
	log.Println("Check URLs by WebRisk API")
	threats, err := wr.LookupURLsContext(ctx, urls)
	urlThreats := make([]URLThreat, 0)
	if err == nil && len(threats) != 0 {
		log.Println("Threats were found by WebRisk API")
		for _, threatArray := range threats {
			for _, threat := range threatArray {
				urlThreats = append(urlThreats, URLThreat{
					Url:    threat.Pattern,
					Threat: threat.String(),
				})
			}
		}
	}
	return urlThreats, err
}
