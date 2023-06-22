package resources

import (
	"context"

	"github.com/arnold-jr/cq-source-bitbucket/client"
	bb "github.com/arnold-jr/cq-source-bitbucket/lib"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "bitbucket_users",
		Resolver:  fetchUsers,
		Transform: transformers.TransformWithStruct(&bb.UserForCQ{}),
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	users, err := c.Bitbucket.GetUsers()
	if err != nil {
		return err
	}

	for _, value := range users {
		res <- value
	}
	return nil
}
