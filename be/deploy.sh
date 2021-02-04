#!/bin/bash

cd pkg/insert
go build insert.go
zip insert.zip insert
aws lambda update-function-code --function-name insert-function --zip-file fileb://insert.zip

cd ../get
go build get.go
zip get.zip get
aws lambda update-function-code --function-name get-function --zip-file fileb://get.zip
