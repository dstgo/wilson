package locale

import (
	"github.com/dstgo/filebox"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
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
// return Group
// return error
func (c Conf) Load() (Group, error) {
	names := filebox.ReadDirFullNames(c.Dir)
	group := make(Group, 5)
	for _, name := range names {
		if !filebox.IsDir(name) {
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
