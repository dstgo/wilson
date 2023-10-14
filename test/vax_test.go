package test

import (
	vax2 "github.com/dstgo/wilson/app/pkg/vax"
	"github.com/dstgo/wilson/app/pkg/vax/is"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) Validate(lang string) error {
	return vax2.Struct(&u, lang,
		vax2.Field(&u.Name, vax2.Required, vax2.RangeLength(1, 10, false)),
		vax2.Field(&u.Email, vax2.Required, is.Email),
		vax2.Field(&u.Age, vax2.Required, vax2.Gte(18), vax2.Lte(100)),
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
