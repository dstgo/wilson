package conf

import "time"

type Config struct {
	DefaultNickName       string
	DefaultPasswordExpire time.Duration
}
