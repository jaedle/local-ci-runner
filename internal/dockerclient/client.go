package dockerclient

import (
	"context"
	"github.com/docker/docker/client"
)

type Client struct {
	client *client.Client
}

func New(ctx context.Context) (*Client, error) {
	cl, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	if _, err := cl.Ping(ctx); err != nil {
		return nil, err
	}

	return &Client{
		client: cl,
	}, nil
}
