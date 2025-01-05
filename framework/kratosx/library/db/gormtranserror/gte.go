package gormtranserror

import "sync"

var (
	_ins GormErrorPlugin
	once sync.Once
)

func NewGlobalGormErrorPlugin(opts ...Option) GormErrorPlugin {
	once.Do(func() {
		_ins = NewGormErrorPlugin(opts...)
		if _ins.options().db != nil {
			err := _ins.options().db.Use(_ins)
			if err != nil {
				panic(err)
			}
		}
	})
	return _ins
}
