package errs

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	// recursive error test
	var testErr = NewI18nError("hello").FallBack("hello")

	sub := testErr.Wrap(errors.New("111"))

	outErr := testErr.Wrap(sub)

	fmt.Println(outErr.Error())
}
