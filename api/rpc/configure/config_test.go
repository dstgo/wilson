package configure

import (
	"testing"
	"time"

	"github.com/dstgo/wilson/framework/kratosx/config"
)

func TestSource_Watch(t *testing.T) {
	source := New("localhost:6082", "Resource", "1025D32F6CA7A2A320FE091B22C5DF3C")
	conf := config.New(source)
	if err := conf.Load(); err != nil {
		t.Error(err.Error())
		return
	}
	conf.Watch("business.storage", func(value config.Value) {
		t.Log("storage 配置变更")
	})

	conf.Watch("business.export", func(value config.Value) {
		t.Log("export 配置变更")
	})
	time.Sleep(100 * time.Second)
}
