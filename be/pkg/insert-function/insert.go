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
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
	fmt.Println(name)
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

func GetLastFreeItemId(acc *[]AccountingItem) int {
	max := 0
	for _, item := range *acc {
		if item.ItemId > max {
			max = item.ItemId
		}
	}
	return max + 1
}

func putItemToDB(dynaClient dynamodbiface.DynamoDBAPI, tableName string, item AccountingItem) error {
	accounts, err := fetchItem(dynaClient, tableName)
	if err != nil {
		fmt.Println(err)
		return err
	}
  input := &dynamodb.PutItemInput{
    Item:     map[string]*dynamodb.AttributeValue{
			"itemid": {
				N: aws.String(strconv.Itoa(GetLastFreeItemId(accounts))),
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

	_, err = dynaClient.PutItem(input)
  if err != nil {
    return errors.New("Error input record")
  }
	return nil
}

func main() {
		openDynamoDB("us-east-2")
    lambda.Start(HandleRequest)
}
