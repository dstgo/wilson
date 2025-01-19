{{ range .Errors }}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Is{{.CamelValue}}(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == {{ .Name }}_{{ .Value }}.String() && e.Code == {{ .HTTPCode }}
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func {{ .CamelValue }}() *errors.Error {
    return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), "{{ .Message }}")
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func {{ .CamelValue }}f(msg string, args ...any) *errors.Error {
    return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf("{{ .Message }}: %s", fmt.Sprintf(msg, args...)))
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func {{ .CamelValue }}Wrap(err error) *errors.Error {
    return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf("{{ .Message }}: %s", err))
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func {{ .CamelValue }}Wrapf(err error, msg string, args ...any) *errors.Error {
    return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf("{{ .Message }}: %s: %s", fmt.Sprintf(msg, args...), err))
}
{{- end }}