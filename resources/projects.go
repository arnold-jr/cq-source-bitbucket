package resources

import (
	"context"

	"github.com/arnold-jr/cq-source-bitbucket/client"
	bb "github.com/arnold-jr/cq-source-bitbucket/lib"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "bitbucket_projects",
		Resolver:  fetchProjects,
		Transform: transformers.TransformWithStruct(&bb.Project{}, transformers.WithPrimaryKeys("UUID")),
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	users, err := c.Bitbucket.GetProjects()
	if err != nil {
		return err
	}

	for _, value := range users {
		res <- value
	}
	return nil
}
