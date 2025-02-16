package dockerclient

import (
	"context"
	"github.com/docker/docker/api/types/container"
)

type RunInput struct {
	Name  string
	Image string
}

func (c *Client) Run(ctx context.Context, in RunInput) error {
	out, err := c.client.ContainerCreate(ctx, &container.Config{
		Image: in.Image,
	}, nil, nil, nil, in.Name)
	if err != nil {
		return err
	}

	return c.client.ContainerStart(ctx, out.ID, container.StartOptions{})
}
