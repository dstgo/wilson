package valx

import (
	"strconv"

	json "github.com/json-iterator/go"
)

func Transform(in any, dst any) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func StrToUint32(in string) uint32 {
	uint32Value, _ := strconv.ParseUint(in, 10, 32)
	return uint32(uint32Value)
}

func StrToInt64(in string) int64 {
	val, _ := strconv.ParseInt(in, 10, 64)
	return val
}
