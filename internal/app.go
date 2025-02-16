package internal

import (
	"context"
	"github.com/jaedle/local-ci-runner/assets"
	"github.com/jaedle/local-ci-runner/internal/dockerclient"
)

const builderImage = "jaedle/local-ci-builder:local"

func New() *App {
	return &App{}
}

type App struct {
	dockerClient *dockerclient.Client
}

func (a *App) Bootstrap(ctx context.Context) error {
	if err := a.init(ctx); err != nil {
		return err
	}

	return a.ensureBuilderIsRunning(ctx)
}

func (a *App) init(ctx context.Context) error {
	client, err := dockerclient.New(ctx)
	if err != nil {
		return err
	}

	a.dockerClient = client
	return nil
}

func (a *App) ensureBuilderIsRunning(ctx context.Context) error {
	if err := a.ensureBuilderImage(ctx); err != nil {
		return err
	}

	return a.ensureBuilder(ctx)
}

func (a *App) ensureBuilderImage(ctx context.Context) error {
	dockerfile, err := assets.GetString("builder/Dockerfile")
	if err != nil {
		return err
	}

	_, err = a.dockerClient.Build(ctx, dockerclient.BuildInput{
		Dockerfile: dockerfile,
		Image:      builderImage,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *App) ensureBuilder(ctx context.Context) error {
	return a.dockerClient.Run(ctx, dockerclient.RunInput{
		Name:  "local-ci-builder",
		Image: builderImage,
	})
}
