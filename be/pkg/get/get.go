package main

import (
	"context"
	"fmt"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const DBTableName = "accounting-db"

var (
    dynaClient dynamodbiface.DynamoDBAPI
)

type AccountingItem struct {
	ItemId int `json:"itemid"`
  Description string `json:"description"`
	Date string `json:"date"`
	Type string `json:"type"`
	Price int `json:"price"`
	Debit int `json:"debit"`
	Kredit int `json:"kredit"`
}

func HandleRequestGet(ctx context.Context, name AccountingItem) (string, error) {
	records, err := fetchItem(dynaClient, DBTableName)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf(err.Error()), err
	}
	return fmt.Sprintf("%+v", *records), nil
}


func fetchItem(dynaClient dynamodbiface.DynamoDBAPI, tableName string) (*[]AccountingItem, error) {
	input := &dynamodb.ScanInput{
    TableName: aws.String(tableName),
  }
  result, err := dynaClient.Scan(input)
  if err != nil {
    return nil, errors.New("Error failed to fetch")
  }

  items := new([]AccountingItem)
  err = dynamodbattribute.UnmarshalListOfMaps(result.Items, items)
  return items, nil
}

func openDynamoDB(region string) {
  awsSession, err := session.NewSession(&aws.Config{
    Region: aws.String(region)},
 )
 if err != nil {
     return
 }
 dynaClient = dynamodb.New(awsSession)
}


func main() {
		openDynamoDB("us-east-2")
    lambda.Start(HandleRequestGet)
}
