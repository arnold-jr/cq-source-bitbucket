package client

import (
	"context"
	"fmt"

	bb "github.com/arnold-jr/cq-source-bitbucket/lib"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
)

type ClientConf struct {
	Workspace string
	Password  string
	Username  string
}

type Client struct {
	Logger    zerolog.Logger
	Bitbucket *bb.Client
}

func (c *Client) ID() string {
	return "Bitbucket"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	bitbucketClient := bb.New()

	return &Client{
		Logger:    logger,
		Bitbucket: bitbucketClient,
	}, nil
}
