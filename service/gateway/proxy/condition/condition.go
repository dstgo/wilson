package condition

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dstgo/wilson/service/gateway/config"
)

type Condition interface {
	Prepare() error
	Judge(*http.Response) bool
}

type byStatusCode struct {
	StatusCode  string
	parsedCodes []int64
}

func (c *byStatusCode) Prepare() error {
	c.parsedCodes = make([]int64, 0)
	parts := strings.Split(c.StatusCode, "-")
	if len(parts) == 0 || len(parts) > 2 {
		return fmt.Errorf("invalid condition %s", c.StatusCode)
	}
	c.parsedCodes = []int64{}
	for _, p := range parts {
		code, err := strconv.ParseInt(p, 10, 16)
		if err != nil {
			return err
		}
		c.parsedCodes = append(c.parsedCodes, code)
	}
	return nil
}

func (c *byStatusCode) Judge(resp *http.Response) bool {
	if len(c.parsedCodes) == 0 {
		return false
	}
	if len(c.parsedCodes) == 1 {
		return int64(resp.StatusCode) == c.parsedCodes[0]
	}
	return (int64(resp.StatusCode) >= c.parsedCodes[0]) &&
		(int64(resp.StatusCode) <= c.parsedCodes[1])
}

type byHeader struct {
	Header *config.Header
	parsed struct {
		name   string
		values map[string]struct{}
	}
}

func (c *byHeader) Judge(resp *http.Response) bool {
	v := resp.Header.Get(c.Header.Name)
	if v == "" {
		return false
	}
	_, ok := c.parsed.values[v]
	return ok
}

func (c *byHeader) Prepare() error {
	c.parsed.name = c.Header.Name
	c.parsed.values = map[string]struct{}{}
	if strings.HasPrefix(c.Header.Value, "[") {
		values, err := parseAsStringList(c.Header.Value)
		if err != nil {
			return err
		}
		for _, v := range values {
			c.parsed.values[v] = struct{}{}
		}
		return nil
	}
	c.parsed.values[c.Header.Value] = struct{}{}
	return nil
}

func parseAsStringList(in string) ([]string, error) {
	var out []string
	if err := json.Unmarshal([]byte(in), &out); err != nil {
		return nil, err
	}
	return out, nil
}

func ParseConditon(in []config.Condition) ([]Condition, error) {
	conditions := make([]Condition, 0, len(in))
	for _, rawCond := range in {
		if rawCond.Header != nil {
			cond := &byHeader{
				Header: rawCond.Header,
			}
			if err := cond.Prepare(); err != nil {
				return nil, err
			}
			conditions = append(conditions, cond)
			continue
		}

		if rawCond.StatusCode != "" {
			cond := &byStatusCode{
				StatusCode: rawCond.StatusCode,
			}
			if err := cond.Prepare(); err != nil {
				return nil, err
			}
			conditions = append(conditions, cond)
			continue
		}
		return nil, fmt.Errorf("unknown condition ")
	}
	return conditions, nil
}

func JudgeConditons(conditions []Condition, resp *http.Response, onEmpty bool) bool {
	if len(conditions) <= 0 {
		return onEmpty
	}
	for _, cond := range conditions {
		if cond.Judge(resp) {
			return true
		}
	}
	return false
}
