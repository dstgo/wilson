package dstparser

import (
	"bytes"
	"text/template"
	"unsafe"

	"github.com/mitchellh/mapstructure"

	"github.com/dstgo/wilson/framework/dontstarve/luax"
)

type LevelOverrideItem struct {
	Name  string `mapstructure:"name"`
	Value any    `mapstructure:"value"`
}

// LevelDataOverrides represents level data overrides information
type LevelDataOverrides struct {
	Id                    string  `mapstructure:"id"`
	Name                  string  `mapstructure:"name"`
	Version               float64 `mapstructure:"version"`
	Desc                  string  `mapstructure:"desc"`
	Location              string  `mapstructure:"location"`
	PlayStyle             string  `mapstructure:"playstyle"`
	HideMiniMap           bool    `mapstructure:"hideminimap"`
	MaxPlayerListPosition float64 `mapstructure:"max_playerlist_position"`
	MinPlayerListPosition float64 `mapstructure:"min_playerlist_position"`
	NumRandomSetPieces    int     `mapstructure:"numrandom_set_pieces"`
	OverrideLevelString   bool    `mapstructure:"override_level_string"`

	// setting
	SettingId   string `mapstructure:"settings_id"`
	SettingName string `mapstructure:"settings_name"`
	SettingDesc string `mapstructure:"settings_desc"`

	// worldgen
	WorldGenId   string `mapstructure:"worldgen_id"`
	WorldGenName string `mapstructure:"worldgen_name"`
	WorldGenDesc string `mapstructure:"worldgen_desc"`

	// meta info
	Overrides           []LevelOverrideItem `mapstructure:"overrides"`
	RandomSetPieces     []string            `mapstructure:"random_set_pieces"`
	RequiredPrefabs     []string            `mapstructure:"required_prefabs"`
	RequiredSetPieces   []string            `mapstructure:"required_setpieces"`
	Substitutes         []string            `mapstructure:"substitutes"`
	BackGroundNodeRange []float64           `mapstructure:"background_node_range"`
}

// ParseLevelDataOverrides parses the leveldataoverrides.lua, returns LevelDataOverrides information
func ParseLevelDataOverrides(luaScript []byte) (LevelDataOverrides, error) {
	l := luax.NewVM()
	defer l.Close()
	if err := l.DoString(unsafe.String(unsafe.SliceData(luaScript), len(luaScript))); err != nil {
		return LevelDataOverrides{}, err
	}

	overrideTable := l.ToTable(-1)
	overrideTableL := luax.LTable(overrideTable)

	var levelDataOverrides LevelDataOverrides

	// basic level data
	levelDataOverrides.Id = overrideTableL.GetString("id")
	levelDataOverrides.Desc = overrideTableL.GetString("desc")
	levelDataOverrides.Version = overrideTableL.GetFloat64("version")
	levelDataOverrides.Name = overrideTableL.GetString("name")
	levelDataOverrides.HideMiniMap = overrideTableL.GetBool("hideminimap")
	levelDataOverrides.Location = overrideTableL.GetString("location")
	levelDataOverrides.MaxPlayerListPosition = overrideTableL.GetFloat64("max_playlist_position")
	levelDataOverrides.MinPlayerListPosition = overrideTableL.GetFloat64("min_playlist_position")
	levelDataOverrides.NumRandomSetPieces = int(overrideTableL.GetInt64("numrandom_set_pieces"))
	levelDataOverrides.OverrideLevelString = overrideTableL.GetBool("override_level_string")
	levelDataOverrides.PlayStyle = overrideTableL.GetString("playstyle")

	// setting
	levelDataOverrides.SettingId = overrideTableL.GetString("settings_id")
	levelDataOverrides.SettingDesc = overrideTableL.GetString("settings_desc")
	levelDataOverrides.SettingName = overrideTableL.GetString("settings_name")

	// world gen
	levelDataOverrides.WorldGenId = overrideTableL.GetString("worldgen_id")
	levelDataOverrides.WorldGenDesc = overrideTableL.GetString("worldgen_desc")
	levelDataOverrides.WorldGenName = overrideTableL.GetString("worldgen_name")

	// world override options
	overrideTableL.GetTable("overrides").MapForEach(func(k string, v luax.Value) {
		levelDataOverrides.Overrides = append(levelDataOverrides.Overrides, LevelOverrideItem{
			Name:  k,
			Value: optionValue(v),
		})
	})

	// random_set_pieces
	overrideTableL.GetTable("random_set_pieces").ArrayForEach(func(index int, value luax.Value) {
		levelDataOverrides.RandomSetPieces = append(levelDataOverrides.RandomSetPieces, value.ToString())
	})

	// random_set_pieces
	overrideTableL.GetTable("required_prefabs").ArrayForEach(func(index int, value luax.Value) {
		levelDataOverrides.RequiredPrefabs = append(levelDataOverrides.RequiredPrefabs, value.ToString())
	})

	// random_set_pieces
	overrideTableL.GetTable("required_setpieces").ArrayForEach(func(index int, value luax.Value) {
		levelDataOverrides.RequiredSetPieces = append(levelDataOverrides.RequiredSetPieces, value.ToString())
	})

	// substitutes
	overrideTableL.GetTable("substitutes").ArrayForEach(func(index int, value luax.Value) {
		levelDataOverrides.Substitutes = append(levelDataOverrides.Substitutes, value.ToString())
	})

	// background_node_range
	overrideTableL.GetTable("background_node_range").ArrayForEach(func(index int, value luax.Value) {
		levelDataOverrides.BackGroundNodeRange = append(levelDataOverrides.BackGroundNodeRange, value.ToFloat64())
	})

	return levelDataOverrides, nil
}

const masterLevelOverrides = `
return {
  desc={{ t .desc }},
  hideminimap={{ t .hideminimap }},
  id={{ t .id }},
  location={{ t .location }},
  max_playlist_position={{ t .max_playlist_position }},
  min_playlist_position={{ t .min_playlist_position }},
  name={{ t .name }},
  numrandom_set_pieces={{ t .numrandom_set_pieces }},
  override_level_string={{ t .override_level_string }},
  overrides={ {{ range $index, $value := .overrides }}
    {{$value.Name }}={{ t $value.Value }}, {{ end }}
  },
  playstyle={{ t .playstyle }},
  random_set_pieces={ {{ range $index, $value := .random_set_pieces }}
    {{ t $value }}, {{ end }}
  },
  required_prefabs={ {{ range $index, $value := .required_prefabs }}
    {{ t $value }}, {{ end }}
  },
  required_setpieces={ {{ range $index, $value := .required_setpieces }}
    {{ t $value }}, {{ end }}
  },
  settings_desc={{ t .settings_desc }},
  settings_id={{ t .settings_id }},
  settings_name={{ t .settings_name }},
  substitutes={ {{ range $index, $value := .substitutes }}
    {{ t $value }}, {{ end }}
  },
  version={{ t .version }},
  worldgen_desc={{ t .worldgen_desc }},
  worldgen_id={{ t .worldgen_id }},
  worldgen_name={{ t .worldgen_name }},
}
`

const caveLevelOverrides = `
return {
  background_node_range={ {{ range $index, $value := .background_node_range }}
    {{ t $value }}, {{ end }}
  },
  desc={{ t .desc }},
  hideminimap={{ t .hideminimap }},
  id={{ t .id }},
  location={{ t .location }},
  max_playlist_position={{ t .max_playlist_position }},
  min_playlist_position={{ t .min_playlist_position }},
  name={{ t .name }},
  numrandom_set_pieces={{ t .numrandom_set_pieces }},
  override_level_string={{ t .override_level_string }},
  overrides={ {{ range $index, $value := .overrides }}
    {{$value.Name }}={{ t $value.Value }}, {{ end }}
  },
  required_prefabs={  {{ range $index, $value := .required_prefabs }}
    {{ t $value }}, {{ end }}
  },
  settings_desc={{ t .settings_desc }},
  settings_id={{ t .settings_id }},
  settings_name={{ t .settings_name }},
  substitutes={ {{ range $index, $value := .required_prefabs }}
    {{ t $value }}, {{ end }}
  },
  version={{ t .version }},
  worldgen_desc={{ t .worldgen_desc }},
  worldgen_id={{ t .worldgen_id }},
  worldgen_name={{ t .worldgen_name }}
}
`

// ToMasterLevelDataOverridesLua converts LevelDataOverrides to lua script
func ToMasterLevelDataOverridesLua(overrides LevelDataOverrides) ([]byte, error) {
	return toLevelDataOverridesLua(masterLevelOverrides, overrides)
}

// ToCaveLevelDataOverridesLua converts LevelDataOverrides to lua script
func ToCaveLevelDataOverridesLua(overrides LevelDataOverrides) ([]byte, error) {
	return toLevelDataOverridesLua(caveLevelOverrides, overrides)
}

func toLevelDataOverridesLua(tmpl string, val any) ([]byte, error) {
	var data map[string]any
	if err := mapstructure.Decode(val, &data); err != nil {
		return nil, err
	}
	templ, err := template.New("leveloverrides").
		Funcs(map[string]any{"t": t}).
		Parse(tmpl)

	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(nil)
	if err := templ.Execute(buffer, data); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
