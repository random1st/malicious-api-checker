package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"os"
)

type URLRecord struct {
	Url string
}

type URLThreat struct {
	Url    string
	Threat string
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func HandleRequest(ctx context.Context, event events.CloudWatchEvent) (string, error) {
	slAPI := createSlackAPI()

	dbAPI, _ := createDynamoDBAPI(ctx)
	urlRecords, err := dbAPI.GetUrls(ctx)

	sb, err := createSafeBrowsingAPI()
	if err == nil {
		sbResult, err := sb.CheckUrls(ctx, urlRecords)
		if err != nil {
			return "", err
		}
		if sbResult == true {
			_ = slAPI.sendMessage("MaliciousAPICheck", "SafeBrowsing check failed")
		}
	}

	wr, err := createWebRiskAPI()
	if err == nil {
		wrResult, err := wr.CheckUrls(ctx, urlRecords)
		if err != nil {
			return "", err
		}
		if len(wrResult) != 0 {
			message := "\n"
			for _, urlThreat := range wrResult {
				message += urlThreat.Url + ":" + urlThreat.Threat + "\n"
			}
			_ = slAPI.sendMessage("WebRiskCheck", "WebRisk check failed"+message)
		}
	}
	return "OK", nil
}

func main() {
	lambda.Start(HandleRequest)
}
