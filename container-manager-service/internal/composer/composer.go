package composer

import (
	"context"
	"io/ioutil"
	"os/exec"
	"time"
)

const (
	dockerCli     = "docker"
	dockerCompose = dockerCli + "-compose"

	up   = "up"
	down = "down"
)

type composer struct {
	cmd exec.Cmd
}

func (c *composer) Compose(ctx context.Context, service string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	dockerComposeFile := c.getDockerComposeFile(service)
	_, err := ioutil.ReadFile(dockerComposeFile)
	if err != nil {
		return err
	}

	c.cmd = *c.buildCommand(ctx, dockerComposeFile)

	return c.cmd.Run()
}

func (c *composer) buildCommand(ctx context.Context, dockerComposeFile string) *exec.Cmd {
	return exec.CommandContext(ctx, dockerCompose, "-f", dockerComposeFile, up)
}

func (c *composer) getDockerComposeFile(service string) string {
	return "compose/" + service + "/docker-compose.yml"
}

func NewComposer() *composer {
	return &composer{}
}
