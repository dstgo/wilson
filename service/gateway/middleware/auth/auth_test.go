package auth

import "testing"

func TestMiddleware(t *testing.T) {
	type item = struct {
		Path   string
		Method string
	}

	object := Auth{
		Whitelist: []item{
			{
				Path:   "/hello/*",
				Method: "POST",
			},
			{
				Path:   "/welcome/lihua",
				Method: "POST",
			},
		},
	}

	tests := []struct {
		input  item
		result bool
	}{
		{
			input: item{
				Path:   "/hello/lihua",
				Method: "POST",
			},
			result: true,
		},
		{
			input: item{
				Path:   "/welcome/lihua1",
				Method: "POST",
			},
			result: false,
		},
		{
			input: item{
				Path:   "/welcome/li",
				Method: "POST",
			},
			result: false,
		},
		{
			input: item{
				Path:   "/hello/",
				Method: "POST",
			},
			result: false,
		},
	}

	for _, item := range tests {
		if object.isWhitelist(item.input.Method, item.input.Path) != item.result {
			t.Error("result error " + item.input.Path)
		}
	}
}
