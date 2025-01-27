package sqlx

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Json is a generic struct for encoding and decoding JSON values,
// implementing the driver.Valuer and sql.Scanner interfaces.
type Json[T any] struct {
	Val T
}

func (j *Json[T]) Marshal() ([]byte, error) {
	return json.Marshal(j.Val)
}

func (j *Json[T]) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &j.Val)
}

func (j *Json[T]) Value() (driver.Value, error) {
	marshal, err := j.Marshal()
	if err != nil {
		return nil, err
	}

	return string(marshal), err
}

func (j *Json[T]) Scan(src any) error {
	if j == nil {
		return nil
	}

	if src == nil {
		return nil
	}

	switch src := src.(type) {
	case string:
		return j.Unmarshal([]byte(src))
	case []byte:
		return j.Unmarshal(src)
	}

	return fmt.Errorf("expected string or []byte, got %T", src)
}
