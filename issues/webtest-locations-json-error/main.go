package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/appinsights/mgmt/insights"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/pkg/errors"
)

func main() {
	const (
		AZURE_SUBSCRIPTION_IS  = auth.SubscriptionID
		AZURE_APPINSIGHTS_NAME = "AZURE_APPINSIGHTS_NAME"
		AZURE_APPINSIGHTS_RG   = "AZURE_APPINSIGHTS_RG"
	)

	ctx := context.TODO()

	subID := os.Getenv(AZURE_SUBSCRIPTION_IS)
	if subID == "" {
		CheckErr(errors.Errorf("env var %s is not defined", AZURE_SUBSCRIPTION_IS))
	}

	aiName := os.Getenv(AZURE_APPINSIGHTS_NAME)
	if aiName == "" {
		log.Printf("Define %s env var with the name of application insights component", AZURE_APPINSIGHTS_NAME)
		CheckErr(errors.Errorf("env var %s is not defined", AZURE_APPINSIGHTS_NAME))
	}

	aiRg := os.Getenv(AZURE_APPINSIGHTS_RG)
	if aiRg == "" {
		log.Printf("Define %s env var with the resource group name of application insights component", AZURE_APPINSIGHTS_RG)
		CheckErr(errors.Errorf("env var %s is not defined", AZURE_APPINSIGHTS_RG))
	}

	authoriser, err := auth.NewAuthorizerFromCLI()
	CheckErr(err)

	locationsClient := insights.NewWebTestLocationsClient(subID)
	locationsClient.Authorizer = authoriser
	_, gotErr := locationsClient.List(ctx, aiRg, aiName)
	var wantErr error = nil

	if gotErr == wantErr {
		log.Printf("PASS: bug fixed")
	} else {
		log.Printf("BUG: webtest-locations-json-error\n")
		log.Printf("  - Want: %v\n", wantErr)
		log.Printf("  - Got : %v\n", gotErr)
		log.Fatalf("BOOM\n")
	}
}

// CheckErr prints a user friendly error to STDERR and exits with a non-zero
// exit code. Unrecognized errors will be printed with an "error: " prefix.
//
func CheckErr(err error) {
	if err == nil {
		return
	}

	// Borrowed from https://github.com/kubernetes/kubectl/blob/master/pkg/cmd/util/helpers.go#L91
	msg := err.Error()
	if len(msg) > 0 {
		msg = fmt.Sprintf("error: %s", err.Error())

		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		_, _ = fmt.Fprint(os.Stderr, msg)
		os.Exit(1)
	}
}
