module github.com/figg/cq-source-bitbucket

go 1.19

require (
	github.com/apache/arrow/go/v13 v13.0.0-20230531201200-cbc17a98dfd9
	github.com/cloudquery/plugin-pb-go v1.0.9
	github.com/cloudquery/plugin-sdk/v3 v3.10.4
	github.com/rs/zerolog v1.29.0
)

replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 v13.0.0-20230606001313-88d5dc2ed455

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/ktrysmt/go-bitbucket v0.9.60 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/oauth2 v0.9.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
