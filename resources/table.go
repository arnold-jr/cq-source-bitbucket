package resources

import (
	"context"
	"fmt"

	"github.com/arnold-jr/cq-source-bitbucket/client"
	"github.com/davecgh/go-spew/spew"
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

	page := 1
	//repo_opts := &bb.RepositoriesOptions{Owner: "figg", Role: "member"}
	repo_opts := &bb.RepositoriesOptions{Owner: "figg", Role: "member", Page: &page}
	resp, err := c.Bitbucket.Repositories.ListForAccount(repo_opts)
	
	spew.Dump(resp)
	spew.Dump(err)
	
	if err != nil {
		fmt.Println(err)	
		return err
	}

	workspaceName := "Figg"
	resp, err := &c.Bitbucket.Workspaces.Get(workspaceName)
	
	spew.Dump(resp)
	spew.Dump(err)
	

	for _, value := range resp.Items {
		res <- value 
		fmt.Println(res)	
	}
	return nil
}
