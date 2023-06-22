package client

import (
	"context"
	"fmt"
	"os"

	//bb "github.com/arnold-jr/cq-source-bitbucket/lib"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
)

type ClientConf struct {
	Workspace string
	Username string
	Password string
}

type Client struct {
	Logger zerolog.Logger
	Config ClientConf
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "Bitbucket"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
  
  getEnv := func(key string) (string, error) {
      val, ok := os.LookupEnv(key)
      if !ok {
          fmt.Printf("%s not set\n", key)
					return "", fmt.Errorf("failed")
      } else {
          fmt.Printf("%s=%s\n", key, val)
          return val, nil
      }
  }
	//workspace, wsNotSet := os.LookupEnv("BITBUCKET_WORKSPACE")
	workspace, _ := getEnv("BITBUCKET_WORKSPACE")	
	//if wsNotSet := true; {
	//	return nil, fmt.Error("BITBUCKET_WORKSPACE not set")
	//}
	bitbucketUser, _ := getEnv("BITBUCKET_USERNAME")
	bitbucketPass, _ := getEnv("BITBUCKET_PASSWORD")

	conf := ClientConf{Workspace: workspace, Username: bitbucketUser, Password: bitbucketPass}
	
	return &Client{
		Logger: logger,
		Config: conf,
	}, nil
}
