package locale

import (
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Conf struct {
	Locale string `mapstructure:"lang"`
	Dir    string `mapstructure:"dir"`
}

// Load read lang file from specified lang dir
// locale name is the language file name without ext
// eg: zh_CN.toml -> zh_CN
func (c Conf) Load() (Group, error) {

	group := make(Group, 5)
	entries, err := os.ReadDir(c.Dir)
	if err != nil {
		return group, err
	}
	// collect all available language-files from lang dir
	for _, entry := range entries {
		if !entry.IsDir() {
			name := filepath.Join(c.Dir, entry.Name())
			v := viper.New()
			v.SetConfigFile(name)
			if err := v.ReadInConfig(); err != nil {
				return nil, err
			}
			localename, _ := strings.CutSuffix(path.Base(name), filepath.Ext(name))
			group[localename] = v
		}
	}

	if len(group) == 0 {
		return nil, errors.Wrap(LocaleFileLeastOneErr, cast.ToString(len(group)))
	}

	return group, nil
}
