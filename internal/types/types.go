package types

var (
	NopObj = NopType{}
)

type (
	H       = map[string]any
	S       = []any
	NopType = struct{}
	Strings = []string
)

// Response
// just used to generate swagger api doc, you should use resp.Response instead
type Response struct {
	Code int    `json:"code" example:"2000"`
	Msg  string `json:"msg" example:"operation success"`
	Err  string `json:"err"`
	Data any    `json:"data"`
}
