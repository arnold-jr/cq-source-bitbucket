package resources

import (
	"context"
	"fmt"
	"os"

	//"github.com/arnold-jr/cq-source-bitbucket/client"
	bb "github.com/arnold-jr/cq-source-bitbucket/lib"
	"github.com/davecgh/go-spew/spew"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Bitbucket() *schema.Table {
	return &schema.Table{
		Name:     "bitbucket_sample_table",
		Resolver: fetchRepos,
		Transform: transformers.TransformWithStruct(&bb.Repository{}),
	}
}

func fetchRepos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	//c := meta.(*client.Client)

	workspace := os.Getenv("BITBUCKET_WORKSPACE")
	bitbucketUser := os.Getenv("BITBUCKET_USERNAME")
	bitbucketPass := os.Getenv("BITBUCKET_PASSWORD")

	repositories, err := bb.GetRepositories(workspace, bitbucketPass, bitbucketUser)
	
	spew.Dump(repositories)
	spew.Dump(err)

	if err != nil {
		return err
	}

	for _, value := range repositories {
		res <- value 
		fmt.Println(res)	
	}
	return nil
}
