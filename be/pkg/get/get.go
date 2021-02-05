package main

import (
	"context"
//	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
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

func HandleRequestGet(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Content-Encoding"] = "application/json"
	records, err := fetchItem(dynaClient, DBTableName)
	if err != nil {
		fmt.Println(err)
		resp.Body = fmt.Sprintf(err.Error())
		resp.StatusCode = 401
		return resp, err
	}
	resp.StatusCode = 200
//	resp.Body = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%+v", records)))
//	resp.IsBase64Encoded = true
	serialized, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err)
		resp.StatusCode = 402
		return resp, err
	}
	resp.Body = string(serialized)
	return resp, nil
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
