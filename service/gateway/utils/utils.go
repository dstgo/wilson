package utils

import (
	"github.com/mitchellh/mapstructure"
)

func Copy(in interface{}, v interface{}) error {
	dc := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           v,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}

	decoder, err := mapstructure.NewDecoder(dc)
	if err != nil {
		return err
	}
	return decoder.Decode(in)
}
