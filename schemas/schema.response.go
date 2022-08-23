package schemas

type ResponseSchema struct {
	StatusCode int
	Method     string
	Message    string
	Data       interface{}
}
