#!/bin/bash

make install

aws lambda update-function-code --function-name insert-function --zip-file fileb://insert.zip --profile be
aws lambda update-function-code --function-name get-function --zip-file fileb://get.zip --profile be

make clean
