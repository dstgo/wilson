package valx

import (
	"regexp"
)

var (
	_PhoneRegex  = regexp.MustCompile(`^1[3456789]\d{9}$`)
	_NumberRegex = regexp.MustCompile(`^-?\d*\.?\d+$`)
	_EmailRegex  = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(.[a-zA-Z]{2,})+$`)
	_IDCardRegex = regexp.MustCompile(`^([1-9]\d{5})(\d{4})(0[1-9]|1[0-2])([0-2][1-9]|[1-3]\d|4[0-6]|5[0-2])(\d{3})(\d|[Xx])$`)
)

func IsEmail(s string) bool {
	return _EmailRegex.MatchString(s)
}

func IsPhone(s string) bool {
	return _PhoneRegex.MatchString(s)
}

func IsIDCard(s string) bool {
	return _IDCardRegex.MatchString(s)
}

func IsNumber(s string) bool {
	return _NumberRegex.MatchString(s)
}
