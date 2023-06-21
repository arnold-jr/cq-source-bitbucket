package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger zerolog.Logger
	Bitbucket *bitbucket.Client
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "Bitbucket"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	bitbucketClientID := os.Getenv("BITBUCKET_USERNAME")
	bitbucketSecret := os.Getenv("BITBUCKET_PASSWORD")
	
	c := bitbucket.NewBasicAuth(bitbucketClientID, bitbucketSecret)	

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
	// TODO: Add your client initialization here

	return &Client{
		Logger: logger,
		Bitbucket: c,
	}, nil
}
