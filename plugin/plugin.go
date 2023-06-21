package plugin

import (
	"github.com/figg/cq-source-bitbucket/client"
	"github.com/figg/cq-source-bitbucket/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"figg-bitbucket",
		Version,
		schema.Tables{
			resources.SampleTable(),
		},
		client.New,
	)
}
