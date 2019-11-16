#!/bin/bash

export AWS_SECRET_ACCESS_KEY=anything
export AWS_ACCESS_KEY_ID=anything
export AWS_REGION=us-east-1

aws secretsmanager create-secret --endpoint-url=http://localhost:4584 --name "DATABASE_CREDS" --description "DB secret" --secret-string '{"user":"dbuser","password":"prod_db_creds"}' --region=us-east-1
