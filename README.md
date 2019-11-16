# AWSSM2env

Provides a convenient way to launch a subprocess with environment variables populated from AWS Secrets Manager.

## Usage

`awssm2env <command> [<arg...>]`

## Application

1. Overwrites matching values for system ENV keys with values stores in AWS Secrets Manager.
2. Launches desired final application acting as a wrapper/entrypoint for easy secrets injection to deployment flow.

## Prerequisites

### 1. Set environment variables for AWS SM connection

```console
export AWS_SECRET_ACCESS_KEY=anything
export AWS_ACCESS_KEY_ID=anything
export AWS_REGION=us-east-1
```

### 2. Credentials stored in AWS Secrets Manager need to follow `awssm2env` retrieval convention

* Desired environment entry:

```console
DATABASE_PASSWORD=prod_db_secret1
```

* AWS SM:

```console
aws --endpoint-url=http://localhost:4584 secretsmanager create-secret --name "DATABASE_PASSWORD" --description "DB secret" --secret-string '{"password":"prod_db_creds"}' --region=us-east-1
```

### 3. Optional environment variables

```console
AWS_ENDPOINT="http://localhost:4584"
AWS_DEBUG=true
```

## TODO

* Integration tests
* Switch initial secrets retrieval from AWS SM API to ListSecrets action (currently not support in Localstack)
* Check signal handling in various use-cases
