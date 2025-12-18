package errors

type errorString struct{
	errs string
}

func New(text string) error {
	return &errorString{errs:text}
}

func (e *errorString) Error() string {
	return e.errs
}