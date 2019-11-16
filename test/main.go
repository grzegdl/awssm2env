package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	t := time.Now()
	databaseUser := os.Getenv("DATABASE_USERNAME")
	databasePass := os.Getenv("DATABASE_PASSWORD")
	fmt.Printf("%s Database User: %s\n", t.Format(time.ANSIC), databaseUser)
	fmt.Printf("%s Database Password: %s\n", t.Format(time.ANSIC), databasePass)

}
