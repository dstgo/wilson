package dstparser

import (
	"bytes"
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
	l := luax.NewVM()
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
	modInfo, err := parseModSimpleInfo(luax.LTable(l.G.Global))
	if err != nil {
		return ModInfo{}, err
	}

	// parse options
	modInfo.ConfigurationOptions = parseModOptions(luax.LTable(l.G.Global).GetTable("configuration_options"))

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
func parseModSimpleInfo(modTable luax.Table) (ModInfo, error) {
	var modInfo ModInfo

	// basic info
	modInfo.Id = modTable.GetString("id")
	modInfo.Name = modTable.GetString("name")
	modInfo.Description = modTable.GetString("description")
	modInfo.Author = modTable.GetString("author")
	modInfo.Version = modTable.GetString("version")

	// dont starve info
	modInfo.ApiVersion = int(modTable.GetInt64("api_version"))
	modInfo.DontStarveCompatible = modTable.GetBool("dont_starve_compatible")
	modInfo.ReignOfGiantsCompatible = modTable.GetBool("reign_of_giants_compatible")
	modInfo.ShipWreckedCompatible = modTable.GetBool("shipwrecked_compatible")
	modInfo.HamletCompatible = modTable.GetBool("hamlet_compatible")

	// dont starve together info
	modInfo.ApiVersionDst = int(modTable.GetInt64("api_version_dst"))
	modInfo.DstCompatible = modTable.GetBool("dst_compatible")
	modInfo.AllClientRequired = modTable.GetBool("all_client_required")
	modInfo.ClientOnly = modTable.GetBool("client_only_mod")
	modInfo.ServerOnly = modTable.GetBool("server_only_mod")
	modInfo.ForgeCompatible = modTable.GetBool("forge_compatible")

	// meta info
	modTable.GetTable("server_filter_tags").ArrayForEach(func(index int, value luax.Value) {
		modInfo.FilterTags = append(modInfo.FilterTags, value.ToString())
	})
	modInfo.Priority = modTable.GetFloat64("priority")
	modInfo.Icon = modTable.GetString("icon")
	modInfo.IconAtlas = modTable.GetString("icon_atlas")

	return modInfo, nil
}

// parse configuration_options from lua script
func parseModOptions(options luax.Table) []ModOption {
	var modOptions []ModOption

	// iterate configuration_options
	options.ArrayForEach(func(index int, value luax.Value) {
		if value.Type() != lua.LTTable {
			return
		}
		var modOption ModOption

		optTable := value.ToTable()

		// option info
		modOption.Name = optTable.GetString("name")
		modOption.Label = optTable.GetString("label")
		modOption.Hover = optTable.GetString("hover")
		modOption.Client = optTable.GetBool("client")

		// default value
		modOption.Default = optionValue(optTable.Get("default"))

		// option items
		modOption.Options = parseModOptionItems(optTable.GetTable("options"))

		optTable.GetTable("tags").ArrayForEach(func(index int, value luax.Value) {
			modOption.Tags = append(modOption.Tags, value.ToString())
		})

		modOptions = append(modOptions, modOption)
	})

	return modOptions
}

// parse mod option items
func parseModOptionItems(optTable luax.Table) []ModOptionItem {
	var items []ModOptionItem
	// iterate items
	optTable.ArrayForEach(func(index int, item luax.Value) {
		if item.Type() != lua.LTTable {
			return
		}

		// build item
		var modItem ModOptionItem

		itemTable := item.ToTable()
		modItem.Description = itemTable.GetString("description")

		dataValue := itemTable.Get("data")
		modItem.Data = optionValue(dataValue)

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
	l := luax.NewVM()
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
		table.GetTable("configuration_options").MapForEach(func(name string, data luax.Value) {
			var item ModOverRideOptionItem
			item.Name = name
			item.Value = optionValue(data)
			items = append(items, item)
		})

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

func optionValue(value luax.Value) any {
	switch value.Type() {
	default:
		return "unknown type"
	case lua.LTString:
		return value.ToString()
	case lua.LTNumber:
		return value.ToFloat64()
	case lua.LTBool:
		return value.ToBool()
	}
}
