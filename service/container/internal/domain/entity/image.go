package entity

import (
	"github.com/dstgo/wilson/framework/pkg/sqlx"
)

type Image struct {
	ID          int64               `json:"id"`
	ImageID     string              `json:"imageId"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Size        int64               `json:"size"`
	Tags        sqlx.Json[[]string] `json:"tags"`
	OS          string              `json:"os"`
	// which the image was created, formatted in RFC 3339 nanoseconds (time.RFC3339Nano).
	BuildAt   string `json:"buildAt"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
