package routes

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	PATCH   Method = "PATCH"
	DELETE  Method = "DELETE"
	OPTIONS Method = "OPTIONS"
)

func (m Method) String() string { return string(m) }
