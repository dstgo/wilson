package errs

import "errors"

func Join(errs ...error) *JoinError {
	return new(JoinError).Join(errs...)
}

type JoinError struct {
	err error
}

func (j *JoinError) Join(errs ...error) *JoinError {
	j.err = errors.Join(append([]error{j.err}, errs...)...)
	return j
}
func (j *JoinError) Err() error {
	return j.err
}
