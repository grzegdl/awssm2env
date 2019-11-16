package api

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"
)

type Session struct {
	*session.Session
}

type Secret struct {
	Credential string
}

var config aws.Config

// NewSession creates a new Session struct
func NewSession() (*Session, error) {

	config = aws.Config{
		// Region:      aws.String("us-east-1"),
		Credentials: credentials.NewEnvCredentials(),
	}

	value, exists := os.LookupEnv("AWS_ENDPOINT")
	if exists && len(value) > 0 {
		config.Endpoint = aws.String(value)
	}

	value, exists = os.LookupEnv("AWS_DEBUG")
	if exists && value == "true" {
		config.LogLevel = aws.LogLevel(aws.LogDebugWithHTTPBody)
	}

	session, err := session.NewSession(&config)

	if err != nil {
		return nil, err
	}

	return &Session{
		Session: session,
	}, nil
}

// GetPasswordBySecretName gets value of key 'password' for given secret name and returns a map of matched secret names [secret_name]=password
// if parameter contains a secret name which doesn't exist, it won't be returned
func (sess *Session) GetPasswordBySecretName(names map[string]string) (map[string]string, error) {
	// secrets to replace *only those which are in secret manager
	secrets := make(map[string]string)
	// DATABASE_PASSWORD=replaceme.DATABASE_CREDS.password
	for k, v := range names {

		values := strings.SplitN(v, ".", 3)
		if len(values) != 3 || values[0] != "replaceme" {
			continue
		}
		secretName := values[1]
		key := values[2]
		output, err := sess.GetAwsSecretValue(secretName)

		if err != nil {
			if aerr, ok := err.(awserr.RequestFailure); ok {
				if aerr.StatusCode() == 404 { // ignore 404
					continue
				}
			}
			return secrets, err
		}

		var data map[string]string

		err = json.Unmarshal([]byte(*output.SecretString), &data)
		if err != nil {
			return secrets, err
		}
		secrets[k] = data[key]
	}

	return secrets, nil
}

// GetAwsSecretList gets secrets list
func (sess *Session) GetAwsSecretList() (*secretsmanager.ListSecretsOutput, error) {
	svc := secretsmanager.New(sess)

	input := &secretsmanager.ListSecretsInput{}

	result, err := svc.ListSecrets(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return nil, errors.Wrapf(err, "Failed to get secret list from AWS Secrets Manager: %s", aerr.Code())
		}

		return nil, err
	}

	return result, nil
}

// GetAwsSecretValue gets secret value by secret name
func (sess *Session) GetAwsSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {

	svc := secretsmanager.New(sess)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(input)

	if err != nil {
		return nil, err
	}

	return result, nil
}
