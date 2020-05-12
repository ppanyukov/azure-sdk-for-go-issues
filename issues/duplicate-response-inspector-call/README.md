# duplicate-response-inspector-call

This is a repro of an issue when trying to get locations available for web test fails
with JSON deserialisation failure.


To run

```
AZURE_SUBSCRIPTION_ID="YOUR_SUB_ID" 
AZURE_APPINSIGHTS_NAME="name of app insights component"
AZURE_APPINSIGHTS_RG="resource group where app insights live"
go run ./main.go
```

Output is like this:

```
2020/05/12 12:31:22 BUG: webtest-locations-json-error
2020/05/12 12:31:22   - Want: <nil>
2020/05/12 12:31:22   - Got : insights.WebTestLocationsClient#List: Failure responding to request: StatusCode=200 -- Original Error: Error occurred unmarshalling JSON - Error = 'json: cannot unmarshal array into Go value of type insights.ApplicationInsightsWebTestLocationsListResult' JSON = '[{"DisplayName":"North Central US","Tag":"us-il-ch1-azr"},{"DisplayName":"West Europe","Tag":"emea-nl-ams-azr"},{"DisplayName":"Southeast Asia","Tag":"apac-sg-sin-azr"},{"DisplayName":"West US","Tag":"us-ca-sjc-azr"},{"DisplayName":"South Central US","Tag":"us-tx-sn1-azr"},{"DisplayName":"East US","Tag":"us-va-ash-azr"},{"DisplayName":"East Asia","Tag":"apac-hk-hkn-azr"},{"DisplayName":"North Europe","Tag":"emea-gb-db3-azr"},{"DisplayName":"Japan East","Tag":"apac-jp-kaw-edge"},{"DisplayName":"Australia East","Tag":"emea-au-syd-edge"},{"DisplayName":"France Central (Formerly France South)","Tag":"emea-ch-zrh-edge"},{"DisplayName":"France Central","Tag":"emea-fr-pra-edge"},{"DisplayName":"UK South","Tag":"emea-ru-msa-edge"},{"DisplayName":"UK West","Tag":"emea-se-sto-edge"},{"DisplayName":"Brazil South","Tag":"latam-br-gru-edge"},{"DisplayName":"Central US","Tag":"us-fl-mia-edge"}]'
2020/05/12 12:31:22 BOOM
```

