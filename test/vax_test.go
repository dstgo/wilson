package test

import (
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/core/vax/is"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) Validate(lang string) error {
	return vax.Struct(&u, lang,
		vax.Field(&u.Name, vax.Required, vax.RangeLength(1, 10, false)),
		vax.Field(&u.Email, vax.Required, is.Email),
		vax.Field(&u.Age, vax.Required, vax.Gte(18), vax.Lte(100)),
	)
}

func TestStruct(t *testing.T) {
	u := User{
		Name:  "ashdlkashjdlajlsdjlasd",
		Email: "a",
		Age:   101,
	}
	err := u.Validate("")
	assert.NotNil(t, err, "validate err not nil")
}
