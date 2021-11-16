package http

import "github.com/Munovv/broblogo/ContainerManagerService/pkg/service"

type Container struct {
	service service.Container
}

func NewContainerHandler(s service.Container) *Container {
	return &Container{
		service: s,
	}
}
