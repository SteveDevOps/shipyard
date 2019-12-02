package clients

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/stretchr/testify/mock"
)

type MockDocker struct {
	mock.Mock
}

func (m *MockDocker) ContainerCreate(
	ctx context.Context,
	config *container.Config,
	hostConfig *container.HostConfig,
	networkingConfig *network.NetworkingConfig,
	containerName string,
) (container.ContainerCreateCreatedBody, error) {

	args := m.Called(ctx, config, hostConfig, networkingConfig, containerName)

	if c, ok := args.Get(0).(container.ContainerCreateCreatedBody); ok {
		return c, args.Error(1)
	}

	return container.ContainerCreateCreatedBody{}, args.Error(1)
}

func (m *MockDocker) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	args := m.Called(ctx, options)

	if cl, ok := args.Get(0).([]types.Container); ok {
		return cl, nil
	}

	return nil, args.Error(1)
}

func (m *MockDocker) ContainerStart(ctx context.Context, ID string, opts types.ContainerStartOptions) error {
	args := m.Called(ctx, ID, opts)

	return args.Error(0)
}

func (m *MockDocker) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error {
	args := m.Called(ctx, containerID, timeout)

	return args.Error(0)
}

func (m *MockDocker) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	args := m.Called(ctx, containerID, options)

	return args.Error(0)
}

func (m *MockDocker) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error) {
	args := m.Called(ctx, options)

	if n, ok := args.Get(0).([]types.NetworkResource); ok {
		return n, args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockDocker) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	args := m.Called(ctx, name, options)

	if n, ok := args.Get(0).(types.NetworkCreateResponse); ok {
		return n, args.Error(1)
	}

	return types.NetworkCreateResponse{}, args.Error(1)
}

func (m *MockDocker) NetworkRemove(ctx context.Context, networkID string) error {
	args := m.Called(ctx, networkID)

	return args.Error(0)
}