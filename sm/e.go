package sm

type BaseError struct {
	status     string
	statusCode int
	method     string
	message    string
}
