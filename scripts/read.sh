#!/bin/bash

export AWS_SECRET_ACCESS_KEY=anything
export AWS_ACCESS_KEY_ID=anything
export AWS_REGION=us-east-1

aws --endpoint-url=http://localhost:4584 secretsmanager get-secret-value --secret-id "DATABASE_CREDS" --region=us-east-1
