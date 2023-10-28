package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
)

type FormDataDescribe struct {
	Values map[string][]string
	Files  map[string][]MultipartFile
}

type MultipartFile struct {
	Filename string
	Size     int64
}

func DescribeFormData(form *multipart.Form) FormDataDescribe {
	var describe FormDataDescribe
	describe.Values = form.Value
	describe.Files = make(map[string][]MultipartFile)

	for key, headers := range form.File {
		var files []MultipartFile
		for _, header := range headers {
			files = append(files, MultipartFile{
				Filename: header.Filename,
				Size:     header.Size,
			})
		}
		describe.Files[key] = files
	}

	return describe
}

func DescribeRawJson(reader io.Reader) ([]byte, error) {
	all, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(nil)
	err = json.Compact(buffer, all)
	return buffer.Bytes(), err
}

func DescribeJson(reader io.Reader) (map[string]any, error) {
	h := make(map[string]any, 20)
	rawJson, err := DescribeRawJson(reader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawJson, &h)
	return h, err
}
