package main

import (
	"fmt"
	"awssm2env/api"
	"awssm2env/env"
	"os"
	"os/exec"
)

const (
	usage = `Usage: awssm2env ENV1 ENV2 -- <command> [<arg...>]
	 

Application:
1. Overwrites matching values for system ENV keys with values stores in AWS Secrets Manager.
2. Launches desired final application acting as a wrapper/entrypoint for easy secrets injection to deployment flow.

Prerequisites:
Set environment variables for AWS SM connection, i.e:
export AWS_SECRET_ACCESS_KEY=anything
export AWS_ACCESS_KEY_ID=anything
export AWS_REGION=us-east-1
`
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}

	updatePassEnvs()

	execCommand(os.Args[1:])
}

func execCommand(args []string) {

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func updatePassEnvs() {
	envs := env.Envs2Map(os.Environ())
	// envs := make(map[string]string)

	// envs["DATABASE_USERNAME"] = "replaceme.DATABASE_CREDS.user"
	// envs["DATABASE_PASSWORD"] = "replaceme.DATABASE_CREDS.password"
	// envs["DATABASE_NAME"] = "replaceme.DATABASE_CREDS.name"

	envsToUpdate, err := getPasswordsForMatchingEnvs(envs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = env.SetEnvs(envsToUpdate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getPasswordsForMatchingEnvs(e map[string]string) (map[string]string, error) {
	session, err := api.NewSession()
	if err != nil {
		return nil, err
	}

	return session.GetPasswordBySecretName(e)
}
