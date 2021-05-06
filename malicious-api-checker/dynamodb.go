package main

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"os"
)

const TABLE = "DYNAMODB_TABLE"

type DynamoDBAPI struct {
	*dynamodb.Client
	Table string
}

var (
	ErrDbTableNameNotFound  = errors.New("no DynamoDB table provided in env vars")
	ErrFetchDataFromDbTable = errors.New("can't fetch data from DynamoDB")
	ErrUnmarshalItems       = errors.New("can't unmarshall item  from DynamoDB")
)

func createDynamoDBAPI(ctx context.Context) (dbAPI *DynamoDBAPI, err error) {
	dynamodbClient := dynamodb.NewFromConfig(getAWSConfig(ctx))
	table, found := os.LookupEnv(TABLE)
	if !found {
		return nil, ErrDbTableNameNotFound
	}
	log.Println("Try to create DynamoDB API")
	return &DynamoDBAPI{dynamodbClient, table}, nil
}

func getAWSConfig(ctx context.Context) aws.Config {
	log.Println("Load AWS config")
	var err error
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Print(err)
	}
	return cfg
}

func (dbAPI *DynamoDBAPI) GetUrls(ctx context.Context) ([]URLRecord, error) {
	log.Println("Try to load urls from DynamoDB ")
	response, err := dbAPI.Scan(ctx, &dynamodb.ScanInput{
		TableName: &dbAPI.Table,
	})
	if err != nil {
		return nil, ErrFetchDataFromDbTable
	}
	var urlRecords []URLRecord
	err = attributevalue.UnmarshalListOfMaps(response.Items, &urlRecords)
	if err != nil {
		return nil, ErrUnmarshalItems
	}
	return urlRecords, nil
}
