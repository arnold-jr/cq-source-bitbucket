package plugin

import (
	"github.com/arnold-jr/cq-source-bitbucket/client"
	"github.com/arnold-jr/cq-source-bitbucket/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"arnold-jr-bitbucket",
		Version,
		schema.Tables{
			resources.BitbucketRepos(),
			resources.BitbucketUsers(),
			resources.BitbucketProjects(),
		},
		client.New,
	)
}
