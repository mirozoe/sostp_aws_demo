package main

import (
	"context"
	"fmt"
	"errors"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
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

func HandleRequest(ctx context.Context, name AccountingItem) (string, error) {
	putItemToDB(dynaClient, DBTableName, name)
	return fmt.Sprintf("Hello!"), nil
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

func putItemToDB(dynaClient dynamodbiface.DynamoDBAPI, tableName string, item AccountingItem) error {
  input := &dynamodb.PutItemInput{
    Item:     map[string]*dynamodb.AttributeValue{
			"itemid": {
				N: aws.String(strconv.Itoa(item.ItemId)),
			},
			"description": {
				S: aws.String(item.Description),
			},
			"date": {
				S: aws.String(item.Date),
			},
			"type": {
				S: aws.String(item.Type),
			},
			"price": {
				N: aws.String(strconv.Itoa(item.Price)),
			},
			"debit": {
				N: aws.String(strconv.Itoa(item.Debit)),
			},
			"kredit": {
				N: aws.String(strconv.Itoa(item.Kredit)),
			},
		},
    TableName: aws.String(tableName),
  }

	_, err := dynaClient.PutItem(input)
  if err != nil {
    return errors.New("Error input record")
  }
	return nil
}

func main() {
		openDynamoDB("us-east-2")
    lambda.Start(HandleRequest)
}
