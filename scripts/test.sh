#!/bin/bash

export AWS_SECRET_ACCESS_KEY=anything
export AWS_ACCESS_KEY_ID=anything
export AWS_REGION=us-east-1
export secret_name=default_secret
export DATABASE_PASSWORD=replaceme


go run ../main.go $1

#aws --endpoint-url=http://localhost:4584 secretsmanager list-secrets --region=us-east-1

