package daemon

import (
	"github.com/docker/docker/client"
	"github.com/dstgo/wilson/internal/data"
)

func NewContainerHandler(docker *client.Client, data *data.DataSource) *ContainerHandler {
	return &ContainerHandler{
		docker: docker,
		data:   data,
	}
}

type ContainerHandler struct {
	docker *client.Client
	data   *data.DataSource
}
