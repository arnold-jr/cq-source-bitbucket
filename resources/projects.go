package resources

import (
	"context"
	"fmt"

	"github.com/arnold-jr/cq-source-bitbucket/client"
	bb "github.com/arnold-jr/cq-source-bitbucket/lib"
	"github.com/davecgh/go-spew/spew"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:     "bitbucket_projects_table",
		Resolver: fetchProjects,
		Transform: transformers.TransformWithStruct(&bb.Project{},transformers.WithPrimaryKeys("ID")),
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	conf := c.Config
	projects, err := bb.GetProjects(conf.Workspace, conf.Password, conf.Username)
	
	spew.Dump(projects)
	spew.Dump(err)

	if err != nil {
		return err
	}

	for _, value := range projects {
		res <- value 
		fmt.Println(res)	
	}
	return nil
}
