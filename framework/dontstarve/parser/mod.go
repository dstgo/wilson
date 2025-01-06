package dstparser

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"
	"unsafe"

	lua "github.com/yuin/gopher-lua"

	"github.com/dstgo/wilson/framework/dontstarve/luax"
)

type ModInfo struct {
	// basic information
	Id          string `mapstructure:"id"`
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Author      string `mapstructure:"author"`
	Version     string `mapstructure:"version"`
	ForumThread string `mapstructure:"forum_thread"`

	// dont starve
	ApiVersion              int  `mapstructure:"api_version"`
	DontStarveCompatible    bool `mapstructure:"dont_starve_compatible"`
	ReignOfGiantsCompatible bool `mapstructure:"reign_of_giants_compatible"`
	ShipWreckedCompatible   bool `mapstructure:"shipwrecked_compatible"`
	HamletCompatible        bool `mapstructure:"hamlet_compatible"`

	// dont starve together
	ApiVersionDst     int  `mapstructure:"api_version_dst"`
	DstCompatible     bool `mapstructure:"dst_compatible"`
	AllClientRequired bool `mapstructure:"all_client_required"`
	ClientOnly        bool `mapstructure:"client_only_mod"`
	ServerOnly        bool `mapstructure:"server_only_mod"`
	ForgeCompatible   bool `mapstructure:"forge_compatible"`

	// meta info
	FilterTags []string `mapstructure:"server_filter_tags"`
	Priority   float64  `mapstructure:"priority"`
	Icon       string   `mapstructure:"icon"`
	IconAtlas  string   `mapstructure:"icon_atlas"`

	// configuration
	ConfigurationOptions []ModOption `mapstructure:"configuration_options"`
}

// ModOption represents a mod option in 'configuration_options'
type ModOption struct {
	// option name, maybe empty
	Name string `mapstructure:"name"`
	// option label, maybe empty
	Label string `mapstructure:"label"`
	// hover tooltip, maybe empty
	Hover string `mapstructure:"hover"`
	// default value of this option
	Default any      `mapstructure:"default"`
	Client  bool     `mapstructure:"client"`
	Tags    []string `mapstructure:"tags"`
	// options available
	Options []ModOptionItem `mapstructure:"options"`
}

type ModOptionItem struct {
	Description string `mapstructure:"description"`
	Data        any    `mapstructure:"data"`
}

type ModOverRideOptionItem struct {
	Name  string `mapstructure:"name"`
	Value any    `mapstructure:"value"`
}

type ModOverRideOption struct {
	Id      string                  `mapstructure:"id"`
	Enabled bool                    `mapstructure:"enabled"`
	Items   []ModOverRideOptionItem `mapstructure:"options"`
}

// ParseModInfo returns the parsed modinfo from lua script
func ParseModInfo(luaScript []byte) (ModInfo, error) {
	return ParseModInfoWithEnv(luaScript, "", "")
}

// ParseModInfoWithEnv parse mod info from lua script with mod environment variables.
func ParseModInfoWithEnv(luaScript []byte, folderName, locale string) (ModInfo, error) {
	l := lua.NewState()
	defer l.Close()

	// prepare mod pre environment
	// see https://forums.kleientertainment.com/forums/topic/150829-game-update-571392/
	l.SetGlobal("locale", lua.LString(locale))
	// dir name like "workshop-1274919201"
	l.SetGlobal("folder_name", lua.LString(folderName))
	// ChooseTranslationTable function will be called in the script,
	// if is needed to translate configuration_options by specific language
	// egs. ChooseTranslationTable(table,[key])
	l.SetGlobal("ChooseTranslationTable", ChooseTranslationTable(l, locale))

	// parse script
	if err := l.DoString(unsafe.String(unsafe.SliceData(luaScript), len(luaScript))); err != nil {
		return ModInfo{}, err
	}

	// parse simple info
	modInfo, err := parseModSimpleInfo(l.G.Global)
	if err != nil {
		return ModInfo{}, err
	}

	// parse options
	modOptions, err := parseModOptions(luax.LTable(l.G.Global).GetTable("configuration_options").T())
	if err != nil {
		return ModInfo{}, err
	}
	modInfo.ConfigurationOptions = modOptions

	return modInfo, nil
}

// ChooseTranslationTable returns *lua.LFunction, this function used to choose the translation table in lua state.
func ChooseTranslationTable(l *lua.LState, locale string) *lua.LFunction {
	return l.NewFunction(func(fnl *lua.LState) int {

		// check first table param
		translationTable := fnl.ToTable(1)
		if translationTable == lua.LNil || translationTable.Len() == 0 {
			return 1
		}

		// Get specific locale table
		target := translationTable.RawGetString(locale)
		if target != lua.LNil {
			fnl.Push(target)
		} else { // or use the default
			fnl.Push(translationTable.RawGetInt(1))
		}

		return 1
	})
}

// parse simple info
func parseModSimpleInfo(table *lua.LTable) (ModInfo, error) {
	var modinfo ModInfo
	if table == lua.LNil {
		return modinfo, errors.New("nil lua global")
	}
	g := luax.LTable(table)

	// basic info
	modinfo.Id = g.GetString("id")
	modinfo.Name = g.GetString("name")
	modinfo.Description = g.GetString("description")
	modinfo.Author = g.GetString("author")
	modinfo.Version = g.GetString("version")

	// ds
	modinfo.ApiVersion = int(g.GetInt64("api_version"))
	modinfo.DontStarveCompatible = g.GetBool("dont_starve_compatible")
	modinfo.ReignOfGiantsCompatible = g.GetBool("reign_of_giants_compatible")
	modinfo.ShipWreckedCompatible = g.GetBool("shipwrecked_compatible")
	modinfo.HamletCompatible = g.GetBool("hamlet_compatible")

	// dst
	modinfo.ApiVersionDst = int(g.GetInt64("api_version_dst"))
	modinfo.DstCompatible = g.GetBool("dst_compatible")
	modinfo.AllClientRequired = g.GetBool("all_client_required")
	modinfo.ClientOnly = g.GetBool("client_only_mod")
	modinfo.ServerOnly = g.GetBool("server_only_mod")
	modinfo.ForgeCompatible = g.GetBool("forge_compatible")

	// meta info
	if g.GetTable("server_filter_tags") != nil {
		g.GetTable("server_filter_tags").T().ForEach(func(key lua.LValue, value lua.LValue) {
			modinfo.FilterTags = append(modinfo.FilterTags, value.String())
		})
	}
	modinfo.Priority = g.GetFloat64("priority")
	modinfo.Icon = g.GetString("icon")
	modinfo.IconAtlas = g.GetString("icon_atlas")

	return modinfo, nil
}

// parse configuration_options from lua script
func parseModOptions(options *lua.LTable) ([]ModOption, error) {
	if options == nil || options == lua.LNil {
		return nil, errors.New("nil configuration_options table")
	}

	var modOptions []ModOption

	// iterate configuration_options
	options.ForEach(func(index lua.LValue, option lua.LValue) {
		if option.Type() != lua.LTTable {
			return
		}
		var modOption ModOption

		optTable := option.(*lua.LTable)
		loptTable := luax.LTable(optTable)

		if t := loptTable.GetTable("options").T(); t != nil || t != lua.LNil {
			modOption.Options = parseModOptionItems(t)
		}
		modOption.Name = loptTable.GetString("name")
		modOption.Label = loptTable.GetString("label")
		modOption.Hover = loptTable.GetString("hover")
		modOption.Client = loptTable.GetBool("client")

		// default value
		defaultValue := luax.LTable(optTable).Get("default")
		modOption.Default = luax.JudgeOptionValue(defaultValue)

		if loptTable.GetTable("tags") != nil {
			loptTable.GetTable("tags").T().ForEach(func(key lua.LValue, value lua.LValue) {
				modOption.Tags = append(modOption.Tags, value.String())
			})
		}

		modOptions = append(modOptions, modOption)
	})

	return modOptions, nil
}

// parse mod option items
func parseModOptionItems(optTable *lua.LTable) []ModOptionItem {
	if optTable == nil || optTable == lua.LNil {
		return nil
	}

	var items []ModOptionItem
	// iterate items
	optTable.ForEach(func(index lua.LValue, item lua.LValue) {
		if item.Type() != lua.LTTable {
			return
		}

		// build item
		var modItem ModOptionItem

		itemTable := luax.LTable(item.(*lua.LTable))
		modItem.Description = itemTable.GetString("description")

		dataValue := itemTable.Get("data")
		modItem.Data = luax.JudgeOptionValue(dataValue)

		// if it has no description, use the string of data
		if len(modItem.Description) == 0 {
			modItem.Description = fmt.Sprintf("%+v", modItem.Data)
		}

		items = append(items, modItem)
	})

	return items
}

// ParseModOverrides returns the mod override options from modoverrides.lua
func ParseModOverrides(luaScript []byte) ([]ModOverRideOption, error) {
	l := lua.NewState()
	defer l.Close()

	if err := l.DoString(unsafe.String(unsafe.SliceData(luaScript), len(luaScript))); err != nil {
		return nil, err
	}
	var options []ModOverRideOption

	overrideTable := l.ToTable(-1)
	// options
	overrideTable.ForEach(func(key lua.LValue, value lua.LValue) {
		if key.Type() != lua.LTString || value == lua.LNil || !strings.Contains(key.String(), "workshop-") {
			return
		}

		var modOverride ModOverRideOption
		table := luax.LTable(value.(*lua.LTable))

		names := strings.Split(key.String(), "-")
		if len(names) > 1 {
			modOverride.Id = names[1]
		}
		modOverride.Enabled = table.GetBool("enabled")

		// items
		var items []ModOverRideOptionItem
		if table.GetTable("configuration_options") != nil {
			table.GetTable("configuration_options").T().ForEach(func(name lua.LValue, data lua.LValue) {
				var item ModOverRideOptionItem
				item.Name = name.String()
				item.Value = luax.JudgeOptionValue(data)
				items = append(items, item)
			})
		}

		modOverride.Items = items
		options = append(options, modOverride)
	})

	return options, nil
}

const modOverrideTmpl = `return {  {{ range $index, $option := . }}
    ["{{ $option.Id }}"] = {
        ["enabled"] = {{ $option.Enabled }},
        ["configuration_options"] = { {{ range $index, $item := $option.Items }}
            ["{{ $item.Name }}"] = {{ t $item.Value }}, {{ end }}
        }
    }, {{ end }}
}`

func t(val any) (any, error) {
	switch val.(type) {
	case string:
		return fmt.Sprintf(`"%s"`, val), nil
	}
	return val, nil
}

// ToModOverrideLua return the lua representation of the modOverride options,
// the format is same as modoverride.lua
func ToModOverrideLua(options []ModOverRideOption) ([]byte, error) {
	templ := template.New("modoverride").Funcs(map[string]any{
		"t": t,
	})

	templ, err := templ.Parse(modOverrideTmpl)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(nil)
	if err := templ.Execute(buffer, options); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
