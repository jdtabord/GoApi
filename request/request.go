package request

//Request interfaz para las diferentes peeticiones
type Request interface {
	Endpoint() string
}
