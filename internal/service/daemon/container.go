package daemon

import (
	"bytes"
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/pkg/strs"
	"github.com/spf13/cast"
	"io"
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

func (c *ContainerHandler) GetContainerLogs(ctx context.Context, cid string, since, until, tail int64, timestamp bool) (string, error) {

	// build options
	opt := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Since:      cast.ToString(since),
		Until:      cast.ToString(until),
		Tail:       cast.ToString(tail),
		Timestamps: timestamp,
	}

	logBody, err := c.docker.ContainerLogs(ctx, cid, opt)
	if err != nil {
		return "", err
	}

	rawlogs, err := io.ReadAll(logBody)
	if err != nil {
		return "", err
	}

	// remove docker steam bytes in per line
	// know more to access https://docs.docker.com/engine/api/v1.44/#tag/Container/operation/ContainerAttach
	lines := bytes.Split(rawlogs, []byte("\n"))
	for i, _ := range lines {
		if len(lines[i]) > 8 {
			lines[i] = lines[i][8:]
		}
	}
	rawlogs = bytes.Join(lines, []byte("\n"))

	return strs.Byte2Str(rawlogs), nil
}
