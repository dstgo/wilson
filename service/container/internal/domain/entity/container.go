package entity

import (
	"github.com/dstgo/wilson/framework/pkg/sqlx"
)

type Container struct {
	ID          int64  `json:"id"`
	InstanceID  int64  `json:"instanceID"`
	ContainerID string `json:"containerId"`
	NodeID      int64  `json:"NodeID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`

	User       string                       `json:"user"`
	Mac        string                       `json:"mac"`
	Image      string                       `json:"image"`
	HostName   string                       `json:"hostname"`
	DomainName string                       `json:"domainName"`
	WorkDir    string                       `json:"workdir"`
	Entrypoint sqlx.Json[[]string]          `json:"entrypoint"`
	CMD        sqlx.Json[[]string]          `json:"cmd"`
	Env        sqlx.Json[map[string]string] `json:"env"`
	Path       string                       `json:"path"`
	Restart    string                       `json:"restart"`
}

type Port struct {
	ID          int64  `json:"id"`
	ContainerID int64  `json:"containerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Port        int    `json:"port"`
	Bind        int    `json:"bind"`
	Protocol    string `json:"protocol"`
	CreatedAt   int64  `json:"createdAt"`
}

type Volume struct {
	ID          int64  `json:"id"`
	ContainerID int64  `json:"containerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Target      string `json:"target"`
	Dir         bool   `json:"dir"`
	RW          bool   `json:"rw"`
	CreatedAt   int64  `json:"createdAt"`
}

type Quota struct {
	ID          int64 `json:"id"`
	ContainerID int64 `json:"containerId"`
	CPU         int64 `json:"cpu"`
	Memory      int64 `json:"memory"`
	Storage     int64 `json:"storage"`
	Network     int64 `json:"network"`
	CreatedAt   int64 `json:"createdAt"`
	UpdatedAt   int64 `json:"updatedAt"`
}
