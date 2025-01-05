package proxy

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/dstgo/wilson/service/gateway/consts"
)

type Response struct {
	Code     int32             `json:"code,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Data     any               `json:"data"`
	Message  string            `json:"message,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	TraceID  string            `json:"traceId,omitempty"`
}

func ResponseFormat(response *http.Response) []byte {
	b, _ := io.ReadAll(response.Body)

	var res any
	if err := json.Unmarshal(b, &res); err != nil {
		return b
	}

	newRes := Response{
		Code:    consts.HTTP_SUCCESS_CODE,
		Message: consts.HTTP_SUCCESS_MESSAGE,
		Reason:  consts.HTTP_SUCCESS_REASON,
		TraceID: response.Header.Get(consts.TRACE_ID),
	}
	// 上游返回error
	m, ok := res.(map[string]any)

	if ok && m["code"] != nil && m["reason"] != nil {
		newRes.Code, _ = m["code"].(int32)
		newRes.Message, _ = m["message"].(string)
		newRes.Metadata, _ = m["metadata"].(map[string]string)
		newRes.Reason, _ = m["reason"].(string)
	} else {
		newRes.Data = res
	}

	b, _ = json.Marshal(newRes)
	return b
}

func GetData(response *http.Response) []byte {
	b, _ := io.ReadAll(response.Body)
	res := Response{}
	if err := json.Unmarshal(b, &res); err != nil {
		return b
	}
	b, _ = json.Marshal(res.Data)
	return b
}
