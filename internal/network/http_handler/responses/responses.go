package responses

const ResOk int = 1000

type ResponseForm struct {
	Code    int
	Message string
	Data    interface{}
}
