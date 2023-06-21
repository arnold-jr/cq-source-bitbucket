package resources

import (
	"context"
	"fmt"

	"github.com/arnold-jr/cq-source-bitbucket/client"
	
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/ktrysmt/go-bitbucket"
)

func Bitbucket() *schema.Table {
	return &schema.Table{
		Name:     "bitbucket_sample_table",
		Resolver: fetchRepos,
		Transform: transformers.TransformWithStruct(&bitbucket.Repository{}),
	}
}

func fetchRepos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*Client) 
	
	client.Bitbucket.Repositories
	//latest, err := client.XKCD.GetLatestComic(ctx)
	//if err != nil {
	//	return err
	//}
	//res <- latest

	//for i := 1; i < latest.Num; i++ {
	//	comic, err := client.XKCD.GetComic(ctx, i)
	//	if err != nil {
	//		return err
	//	}
	//	res <- comic
	//}
	return nil
}
