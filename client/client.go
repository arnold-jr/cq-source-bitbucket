package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
)

type ClientConf struct {
	Workspace string
	Password string
	Username string
}

type Client struct {
	Logger zerolog.Logger
	Config ClientConf
}

func (c *Client) ID() string {
	return "Bitbucket"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
  
	workspace := os.Getenv("BITBUCKET_WORKSPACE")	
	bitbucketUser := os.Getenv("BITBUCKET_USERNAME")
	bitbucketPass := os.Getenv("BITBUCKET_PASSWORD")

	conf := ClientConf{Workspace: workspace, Password: bitbucketPass, Username: bitbucketUser}
	
	return &Client{
		Logger: logger,
		Config: conf,
	}, nil
}
