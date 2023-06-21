package resources

import (
	"context"

	"github.com/arnold-jr/cq-source-bitbucket/client"
	
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	bb "github.com/ktrysmt/go-bitbucket"
)

func Bitbucket() *schema.Table {
	return &schema.Table{
		Name:     "bitbucket_sample_table",
		Resolver: fetchRepos,
		Transform: transformers.TransformWithStruct(&bb.Repository{}),
	}
}

func fetchRepos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client) 

	foo := "foo"	
	page := 0	
	repo_opts := &bb.RepositoriesOptions{Owner: "", Role: "", Page: &page, Keyword: &foo}
	resp, err := c.Bitbucket.Repositories.ListForAccount(repo_opts)
	if err != nil {
		return err
	}

	for _, value := range resp.Items {
		res <- value 
	}
	return nil
}
