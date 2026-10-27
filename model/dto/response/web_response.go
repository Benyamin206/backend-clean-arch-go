package response

type WebResponse struct {
	Status string
	Code   int
	Data   interface{}
}
