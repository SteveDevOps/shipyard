package clients

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// Docker defines an interface for a Docker client
type Docker interface {
	ContainerCreate(
		ctx context.Context,
		config *container.Config,
		hostConfig *container.HostConfig,
		networkingConfig *network.NetworkingConfig,
		containerName string,
	) (container.ContainerCreateCreatedBody, error)

	ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	ContainerStart(context.Context, string, types.ContainerStartOptions) error
	ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
	ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error

	NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error)
	NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)
	NetworkRemove(ctx context.Context, networkID string) error
}

// NewDocker creates a new Docker client
func NewDocker() (Docker, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return cli, nil
}