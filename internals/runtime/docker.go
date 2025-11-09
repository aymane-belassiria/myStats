package runtime

import (
	"context"
	"fmt"

	"github.com/moby/moby/client"
)

type Container struct {
	ID      string `json:"Id"`
	Name    string
	Image   string
	ImageID string
}

type DockerRuntime struct {
	cli *client.Client
}

func NewDockerRuntime() (*DockerRuntime, error) {
	cli, err := client.New(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &DockerRuntime{cli: cli}, nil
}

func (d *DockerRuntime) ListContainers() ([]Container, error) {
	containers, err := d.cli.ContainerList(context.Background(), client.ContainerListOptions{All: false})
	if err != nil {
		return nil, err
	}
	result := []Container{}
	for _, ctr := range containers.Items {
		result = append(result, Container{
			ID:      ctr.ID,
			Name:    ctr.Names[0][1:],
			Image:   ctr.Image,
			ImageID: ctr.ImageID,
		})
	}

	return result, nil
}

func (d *DockerRuntime) GetContainerByPID(pid string) (*Container, error) {
	container, err := d.cli.ContainerInspect(context.Background(), pid, client.ContainerInspectOptions{Size: false})
	if err != nil {
		return nil, err
	}
	if container.Container.State == nil {
		return nil, fmt.Errorf("Container with pid %s has no state info", pid)
	}
	info := &Container{
		ID:      container.Container.ID,
		Name:    container.Container.Name[1:],
		Image:   container.Container.Config.Image,
		ImageID: container.Container.Image,
	}
	return info, nil
}
