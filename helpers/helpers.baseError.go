package helpers

type error interface {
	Error() string
}

type baseError struct {
	status     string
	statusCode int
	method     string
	message    string
}

func BaseError() {

	return
}
