package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeBody(body io.ReadCloser, dto_request interface{}) {
	decoder := json.NewDecoder(body)

	err := decoder.Decode(dto_request)

	PanicIfError(err)
}

func ToResponseWriter(writer http.ResponseWriter, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)

	encoder.Encode(data)
}
