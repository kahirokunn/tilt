package docker

import (
	"context"
	"io"

	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"golang.org/x/sync/errgroup"

	"github.com/tilt-dev/tilt/internal/container"
	"github.com/tilt-dev/tilt/pkg/model"
)

// A docker client that returns errors on every method call.
// Useful when the client failed to init, but we don't have enough
// info at init time to tell if anyone is going to use it.

type explodingClient struct {
	err error
}

func newExplodingClient(err error) explodingClient {
	return explodingClient{err: err}
}

func (c explodingClient) SetOrchestrator(orc model.Orchestrator) {
}
func (c explodingClient) ForOrchestrator(orc model.Orchestrator) Client {
	return c
}
func (c explodingClient) CheckConnected() error {
	return c.err
}
func (c explodingClient) Env() Env {
	return Env{}
}
func (c explodingClient) BuilderVersion() types.BuilderVersion {
	return types.BuilderVersion("")
}
func (c explodingClient) ServerVersion() types.Version {
	return types.Version{}
}
func (c explodingClient) ContainerLogs(ctx context.Context, containerID string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	return nil, c.err
}
func (c explodingClient) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	return types.ContainerJSON{}, c.err
}
func (c explodingClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return nil, c.err
}
func (c explodingClient) ContainerRestartNoWait(ctx context.Context, containerID string) error {
	return c.err
}
func (c explodingClient) Run(ctx context.Context, opts RunConfig) (RunResult, error) {
	return RunResult{}, c.err
}
func (c explodingClient) ExecInContainer(ctx context.Context, cID container.ID, cmd model.Cmd, in io.Reader, out io.Writer) error {
	return c.err
}
func (c explodingClient) ImagePull(_ context.Context, _ reference.Named) (reference.Canonical, error) {
	return nil, c.err
}
func (c explodingClient) ImagePush(ctx context.Context, ref reference.NamedTagged) (io.ReadCloser, error) {
	return nil, c.err
}
func (c explodingClient) ImageBuild(ctx context.Context, g *errgroup.Group, buildContext io.Reader, options BuildOptions) (types.ImageBuildResponse, error) {
	return types.ImageBuildResponse{}, c.err
}
func (c explodingClient) ImageTag(ctx context.Context, source, target string) error {
	return c.err
}
func (c explodingClient) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error) {
	return types.ImageInspect{}, nil, c.err
}
func (c explodingClient) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	return nil, c.err
}
func (c explodingClient) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error) {
	return nil, c.err
}
func (c explodingClient) NewVersionError(apiRequired, feature string) error {
	return c.err
}
func (c explodingClient) BuildCachePrune(ctx context.Context, opts types.BuildCachePruneOptions) (*types.BuildCachePruneReport, error) {
	return nil, c.err
}
func (c explodingClient) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error) {
	return types.ContainersPruneReport{}, c.err
}

var _ Client = &explodingClient{}
