package errs

func IsLocaleError(err error) bool {
	_, ok := err.(LocaleError)
	return ok
}

func NewError() LocaleError {
	return LocaleError{}
}

func NewI18nError(i18n string) LocaleError {
	return NewError().I18n(i18n)
}

// LocaleError
// a response error wrap
// Err field
type LocaleError struct {
	ErrorCode  int
	HttpStatus int
	LangCode   string
	Fb         string
	Err        error
}

func (e LocaleError) Code(code int) LocaleError {
	e.ErrorCode = code
	return e
}

func (e LocaleError) Status(status int) LocaleError {
	e.HttpStatus = status
	return e
}

func (e LocaleError) I18n(langCode string) LocaleError {
	e.LangCode = langCode
	return e
}

func (e LocaleError) FallBack(fallback string) LocaleError {
	e.Fb = fallback
	return e
}

func (e LocaleError) Wrap(err error) LocaleError {
	e.Err = err
	return e
}

func (e LocaleError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Fb
}
