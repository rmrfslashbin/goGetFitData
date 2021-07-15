// Package: main
// Language: go
// Path: main.go
package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/rmrfslashbin/getFitData/googleAuth"
	"golang.org/x/oauth2/google"
	fitness "google.golang.org/api/fitness/v1"
	"google.golang.org/api/option"
)

func main() {
	// Parse command line flags
	var tokenFile = flag.String("f", "token-for-google.json", "FQPN of token file")
	flag.Parse()

	// Read token from file
	b, err := ioutil.ReadFile(*tokenFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, fitness.FitnessActivityReadScope,
		fitness.FitnessBodyReadScope,
		fitness.FitnessLocationReadScope,
	)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	// Create an authorized client and set up a context.
	client := googleAuth.GetClient(config, *tokenFile)
	ctx := context.Background()

	// Set up a new service.
	srv, err := fitness.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Get data.
	ds, err := srv.Users.DataSources.List("me").Do()
	if err != nil {
		log.Fatalf("Unable to fetch: %v", err)
	}

	// Print the data.
	for _, d := range ds.DataSource {
		fmt.Println((d))
	}
}
