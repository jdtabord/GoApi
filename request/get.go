package request

//Get estructura para la peticion get
type Get struct {
	RequestType string
}

//Endpoint para obtener el endpoint
func (get Get) Endpoint() string {
	return get.RequestType
}
