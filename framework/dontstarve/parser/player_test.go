package dstparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseServerChatLogs(t *testing.T) {
	const chatlog = `[00:01:18]: [Join Announcement] 寒江蓑笠翁
[00:01:57]: [Say] (KU_iJIpcpXi) 寒江蓑笠翁: aka
[00:02:53]: [Say] (KU_iJIpcpXi) 寒江蓑笠翁: nihao
[00:05:32]: [Death Announcement] 寒江蓑笠翁 死于： 恶作剧。他变成了可怕的鬼魂！
[00:06:37]: [Resurrect Announcement] 寒江蓑笠翁 复活自： 绚丽之门.
[00:08:09]: [Say] (KU_iJIpcpXi) 寒江蓑笠翁: \n
[00:53:34]: [Say] (KU_iJIpcpXi) 寒江蓑笠翁: (aaa)
[01:35:12]: [Death Announcement] 寒江蓑笠翁 死于： 黑暗。他变成了可怕的鬼魂！
[01:35:20]: [Roll Announcement] (KU_iJIpcpXi) 寒江蓑笠翁 83 (1-100)
[01:35:27]: [Vote Announcement] rollback passed
[01:36:18]: [Say] (KU_iJIpcpXi) 寒江蓑笠翁: dasdasd\nasdajsldjlasjdlkjalskdjalksdljasldjalksjddddddddjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj`

	logs, err := ParseServerChatLogs([]byte(chatlog))
	assert.Nil(t, err)
	t.Log(logs)
}
