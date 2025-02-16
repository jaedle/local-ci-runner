package dockerclient

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"io"
)

type BuildInput struct {
	Dockerfile string
	Image      string
}

type BuildOutput struct {
	ImageID string
}

func (c *Client) Build(ctx context.Context, input BuildInput) (*BuildOutput, error) {
	bfr, err := buildContext(input)
	if err != nil {
		return nil, err
	}

	build, err := c.client.ImageBuild(ctx, bfr, types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{input.Image},
	})

	if err != nil {
		return nil, err
	}

	defer func() { _ = build.Body.Close() }()

	if _, err := io.ReadAll(build.Body); err != nil {
		return nil, err
	}

	if list, err := c.client.ImageList(ctx, image.ListOptions{
		Filters: filters.NewArgs(filters.Arg("reference", input.Image)),
	}); err != nil {
		return nil, err
	} else if len(list) == 0 {
		return nil, errors.New("no images found")
	} else {
		return &BuildOutput{ImageID: list[0].ID}, nil
	}

}

func buildContext(input BuildInput) (*bytes.Buffer, error) {
	result := new(bytes.Buffer)

	gw := gzip.NewWriter(result)
	tw := tar.NewWriter(gw)

	if err := tw.WriteHeader(&tar.Header{
		Name: "Dockerfile",
		Size: int64(len(input.Dockerfile)),
	}); err != nil {
		return nil, err
	}

	if _, err := tw.Write([]byte(input.Dockerfile)); err != nil {
		return nil, err
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}
	if err := gw.Close(); err != nil {
		return nil, err
	}

	return result, nil
}
