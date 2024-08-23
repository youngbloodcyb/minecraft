package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func CreateMinecraftServer(ctx context.Context, containerName string) (string, error) {
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return "", err
    }

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: "itzg/minecraft-server",
        ExposedPorts: nat.PortSet{
            "25565/tcp": struct{}{},
        },
    }, &container.HostConfig{
        PortBindings: nat.PortMap{
            "25565/tcp": []nat.PortBinding{
                {
                    HostIP:   "0.0.0.0",
                    HostPort: "25565",
                },
            },
        },
    }, nil, nil, containerName)
    if err != nil {
        return "", err
    }

    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        return "", err
    }

    return resp.ID, nil
}
