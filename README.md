# AzureDuplicateResponder

This is a repro for an issue which calls response inspector twice in Azure SDK for Go.

It follows the "Writing Custom Request/Response Inspectors" example here: https://github.com/Azure/azure-sdk-for-go#writing-custom-requestresponse-inspectors

To run

```
export AZURE_SUBSCRIPTION_ID="YOUR_SUB_ID"
go run ./main.go
```

Output is like this:

```
2020/05/11 17:43:44 ResponderFunc is called   <== once
2020/05/11 17:43:44 ResponderFunc is called   <== twice
2020/05/11 17:43:44 Got Subscription: YOUR_SUB_NAME
```

It looks like both autorest *and* Azure SDK for Go call the inspector thus resulting in duplication.
