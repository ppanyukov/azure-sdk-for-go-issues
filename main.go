package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/subscriptions"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
)

func main() {
	subID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subID == "" {
		CheckErr(errors.New("env var AZURE_SUBSCRIPTION_ID is not defined"))
	}

	authoriser, err := auth.NewAuthorizerFromCLI()
	CheckErr(err)

	ctx := context.TODO()
	subClient := subscriptions.NewClient()
	subClient.Authorizer = authoriser
	subClient.ResponseInspector = LogResponse()
	result, err := subClient.Get(ctx, subID)
	CheckErr(err)

	log.Printf("Got Subscription: %s\n", to.String(result.DisplayName))
}

func LogResponse() autorest.RespondDecorator {
	return func(p autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(r *http.Response) error {
			err := p.Respond(r)
			CheckErr(err)
			//BUG: this responder will be called twice.
			log.Printf("ResponderFunc is called")
			return err
		})
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
