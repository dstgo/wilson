package coco

import "github.com/gin-gonic/gin"

type GinMode = string

const (
	ReleaseMode GinMode = "release"
	DebugMode           = "debug"
	Test                = "test"
)

func Mode() GinMode {
	return gin.Mode()
}

// SetMode it should be set before httpServer starting
// param mode GinMode
func SetMode(mode GinMode) {
	gin.SetMode(mode)
}
