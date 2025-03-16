package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	json "github.com/json-iterator/go"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
)

type TemplateValue struct {
	Value   string
	Exclude bool
}

type templateParser struct {
}

type TemplateParser interface {
	// CheckTemplate 检查模板变量
	CheckTemplate(keys []string, template string) error

	// RenderTemplate 渲染模板
	RenderTemplate(bvs []*entity.BusinessValue, rvs []*entity.ResourceValue, format string, content string) (string, error)
}

func NewTemplateParser() TemplateParser {
	return &templateParser{}
}

func (t *templateParser) fillKey(val string) string {
	return fmt.Sprintf(`${%s}`, val)
}

// CheckTemplate 校验数据模板数据是否合法
func (t *templateParser) CheckTemplate(keys []string, template string) error {
	// 组合模板和模板的key
	bucket := map[string]bool{}
	for _, key := range keys {
		bucket[t.fillKey(key)] = true
	}

	// 进行增则匹配
	reg := regexp.MustCompile(`\$\{(\w|\.)+}`)
	tempKeys := reg.FindAllString(template, -1)
	// 进行参数判断
	for _, key := range tempKeys {
		if !bucket[key] {
			return errors.New("invalid field：" + key)
		}
	}
	return nil
}

func (t *templateParser) replaceByFormatJson(template, key string, val TemplateValue) string {
	isQuote := strings.Contains(template, fmt.Sprintf(`"%v"`, key)) && strings.Contains(template, fmt.Sprintf(`'%v'`, key))
	if isQuote && val.Exclude {
		template = strings.Replace(template, fmt.Sprintf(`"%v"`, key), val.Value, 1)
		template = strings.Replace(template, fmt.Sprintf(`'%v'`, key), val.Value, 1)
	} else {
		template = strings.Replace(template, key, fmt.Sprintf("%v", val.Value), 1)
	}
	return template
}

func (t *templateParser) replaceByFormatYaml(template, key string, val TemplateValue) string {
	return strings.Replace(template, key, fmt.Sprintf("%v", val.Value), 1)
}

func (t *templateParser) parseBusinessValue(tp, input string) TemplateValue {
	switch tp {
	case "string":
		return TemplateValue{Value: input, Exclude: false}
	case "object":
		return TemplateValue{Value: input, Exclude: true}
	default:
		return TemplateValue{Value: input, Exclude: true}
	}
}

func (t *templateParser) parseResourceValue(val any) TemplateValue {
	if val == nil {
		return TemplateValue{Value: "null", Exclude: true}
	}
	switch val := val.(type) {
	case string:
		return TemplateValue{Value: val, Exclude: false}
	case map[string]any, []any:
		str, _ := json.Marshal(val)
		return TemplateValue{Value: string(str), Exclude: true}
	default:
		return TemplateValue{Value: fmt.Sprint(val), Exclude: true}
	}
}

func (t *templateParser) getVariableValue(bvs []*entity.BusinessValue, rvs []*entity.ResourceValue) (map[string]TemplateValue, error) {
	result := make(map[string]TemplateValue)
	// 解析业务字段
	for _, item := range bvs {
		result[t.fillKey(item.Business.Keyword)] = t.parseBusinessValue(item.Business.Type, item.Value)
	}

	// 解析资源字段
	for _, item := range rvs {
		fs := map[string]any{}
		_ = json.Unmarshal([]byte(item.Value), &fs)

		fields := strings.Split(item.Resource.Fields, ",")
		for _, key := range fields {
			k := fmt.Sprintf("%s.%s", item.Resource.Keyword, key)
			result[t.fillKey(k)] = t.parseResourceValue(fs[key])
		}
	}

	return result, nil
}

// RenderTemplate 指定content生成配置预览
func (t *templateParser) RenderTemplate(bvs []*entity.BusinessValue, rvs []*entity.ResourceValue, format string, content string) (string, error) {
	// 获取字段的配置值
	values, err := t.getVariableValue(bvs, rvs)
	if err != nil {
		return "", err
	}

	// 进行值替换
	reg := regexp.MustCompile(`\$\{(\w|\.)+}`)
	tempKeys := reg.FindAllString(content, -1)
	for _, key := range tempKeys {
		if val, ok := values[key]; !ok {
			return "", errors.New("invalid field：" + key)
		} else {
			if format == "json" {
				content = t.replaceByFormatJson(content, key, val)
			}
			if format == "yaml" {
				content = t.replaceByFormatYaml(content, key, val)
			}
		}
	}
	return content, nil
}
