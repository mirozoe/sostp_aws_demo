package main

import (
	"context"
	"time"
	"errors"
	"fmt"
	"sort"

//	"github.com/aws/aws-lambda-go/events"
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

//func HandleRequestGet(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
//	resp.Headers["Access-Control-Allow-Origin"] = "*"
//	resp.Headers["Content-Encoding"] = "application/json"
//	records, err := fetchItem(dynaClient, DBTableName)
//	if err != nil {
//		fmt.Println(err)
//		resp.Body = fmt.Sprintf(err.Error())
//		resp.StatusCode = 401
//		return resp, err
//	}
//	resp.StatusCode = 200
//	recordsSorted := SortRecordsByDate(*records)
//	serialized, err := json.Marshal(recordsSorted)
//	if err != nil {
//		fmt.Println(err)
//		resp.StatusCode = 402
//		return resp, err
//	}
//	resp.Body = string(serialized)
//	return resp, nil
//	return serialized
//}

func HandleRequestGet(ctx context.Context) ([]AccountingItem, error) {
	records, err := fetchItem(dynaClient, DBTableName)
	if err != nil {
		fmt.Println(err)
		return []AccountingItem{}, err
	}
	recordsSorted := SortRecordsByDate(*records)
	return recordsSorted, nil
}

func parseDate(date string) time.Time {
	t, err := time.Parse("2006-01-02", date) 
	if err != nil {
		fmt.Printf("Cannot parse time %s\n", date)
		return time.Now()
	}
	return t
}

func SortRecordsByDate(records []AccountingItem) []AccountingItem {
	sort.Slice(records, func (i,j int) bool {
			t1 := parseDate(records[i].Date)
			t2 := parseDate(records[j].Date)
			return t1.Before(t2)
	})
	return records
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
