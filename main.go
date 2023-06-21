package main

import (
	"github.com/arnold-jr/cq-source-bitbucket/plugin"

	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
