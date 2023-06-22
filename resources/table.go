package resources

import (
	"context"

	"github.com/arnold-jr/cq-source-bitbucket/client"
	bb "github.com/arnold-jr/cq-source-bitbucket/lib"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Bitbucket() *schema.Table {
	return &schema.Table{
		Name:     "bitbucket_repos_table",
		Resolver: fetchRepos,
		Transform: transformers.TransformWithStruct(&bb.Repository{}),
	}
}

func fetchRepos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	conf := c.Config
	repositories, err := bb.GetRepositories(conf.Workspace, conf.Password, conf.Username)
	
	if err != nil {
		return err
	}

	for _, value := range repositories {
		res <- value 
	}
	return nil
}
