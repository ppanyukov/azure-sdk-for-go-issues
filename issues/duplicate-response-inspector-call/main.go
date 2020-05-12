package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/subscriptions"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/pkg/errors"
)

func main() {
	subID := os.Getenv(auth.SubscriptionID)
	if subID == "" {
		CheckErr(errors.Errorf("env var %s is not defined", auth.SubscriptionID))
	}

	authoriser, err := auth.NewAuthorizerFromCLI()
	CheckErr(err)

	const wantCounter = 1
	gotCounter := 0

	ctx := context.TODO()
	subClient := subscriptions.NewClient()
	subClient.Authorizer = authoriser
	subClient.ResponseInspector = LogResponse(&gotCounter)
	_, err = subClient.Get(ctx, subID)
	CheckErr(err)

	if gotCounter == wantCounter {
		log.Printf("PASS: bug fixed")
	} else {
		log.Printf("BUG: duplicate-response-inspector-call\n")
		log.Printf("  - Want: %d\n", wantCounter)
		log.Printf("  - Got : %d\n", gotCounter)
		log.Fatalf("BOOM\n")
	}
}

func LogResponse(counter *int) autorest.RespondDecorator {
	return func(p autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(r *http.Response) error {
			err := p.Respond(r)
			CheckErr(err)
			//BUG: this responder will be called twice.
			*counter++
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
