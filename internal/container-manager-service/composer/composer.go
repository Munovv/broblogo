package composer

import (
	"os"
	"os/exec"
	"time"
)

const (
	dockerCli     = "docker"
	dockerCompose = dockerCli + "-compose"

	up   = "up"
	down = "stop"
)

type composer struct {
	cmd exec.Cmd
}

func (c *composer) Compose(services []string) error {
	var err error

	c.cmd = *c.upCommand(services)

	go func() {
		c.cmd.Stdout = os.Stdout
		c.cmd.Stderr = os.Stderr
		err = c.cmd.Run()
	}()

	time.Sleep(4 * time.Second)

	return err
}

func (c *composer) Down(services []string) error {
	var err error

	c.cmd = *c.downCommand(services)

	go func() {
		c.cmd.Stdout = os.Stdout
		c.cmd.Stderr = os.Stderr
		err = c.cmd.Run()
	}()

	time.Sleep(4 * time.Second)

	return err
}

func (c *composer) upCommand(services []string) *exec.Cmd {
	baseArgs := []string{"-f", "./deploy/docker-compose.yml", up}
	args := append(baseArgs, append(services)...)

	return exec.Command(dockerCompose, args...)
}

func (c *composer) downCommand(services []string) *exec.Cmd {
	baseArgs := []string{"-f", "./deploy/docker-compose.yml", down}
	args := append(baseArgs, append(services)...)

	return exec.Command(dockerCompose, args...)
}

func NewComposer() *composer {
	return &composer{}
}
