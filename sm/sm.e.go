package sm

// type error interface {
// 	Error() string
// }

type BaseError struct {
	status     string
	statusCode int
	method     string
	message    string
}
