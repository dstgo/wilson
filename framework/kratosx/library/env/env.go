package env

import (
	"os"
)

const (
	appName    = "APP_NAME"
	appVersion = "APP_VERSION"
)

func GetAppName() string {
	return os.Getenv(appName)
}

func SetAppName(val string) {
	_ = os.Setenv(appName, val)
}

func GetAppVersion() string {
	return os.Getenv(appVersion)
}

func SetAppVersion(val string) {
	_ = os.Setenv(appVersion, val)
}

func GetHostName() string {
	id, _ := os.Hostname()
	return id
}
