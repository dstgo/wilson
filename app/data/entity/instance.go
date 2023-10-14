package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Instance represents an instance, usually a docker container
type Instance struct {
	ID    string `gorm:"comment:docker container id;"`
	Uid   string `gorm:"comment:safe unique id, sha1 from instance.id;"`
	Name  string `gorm:"comment:docker container name;"`
	Image string `gorm:"comment:docker image name;"`
	Note  string `gorm:"comment:remark note;"`

	Meta InstanceMeta `gorm:"instance metadata info, json format stored in db"`

	ExpiredAt time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
}

type InstanceMeta map[string]any

func (m InstanceMeta) Scan(src any) error {
	if s, ok := src.(string); ok {
		err := json.Unmarshal([]byte(s), &m)
		if err != nil {
			return err
		}
	}
	return errors.New("can not convert to InstanceMeta")
}

func (m InstanceMeta) Value() (driver.Value, error) {
	marshal, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
