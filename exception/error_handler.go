package exception

import (
	"net/http"

	"github.com/Benyamin206/backend-clean-arch-go/helper"
	"github.com/Benyamin206/backend-clean-arch-go/model/dto/response"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, e interface{}) {

	if notFoundError(w, r, e) {
		return
	}

	internalServerError(w, r, e)
}

func notFoundError(w http.ResponseWriter, r *http.Request, e interface{}) bool {
	exception, ok := e.(NotFoundError)

	if ok {
		webResponse := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: "DATA NOT FOUND",
			Data:   exception.Error,
		}

		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json")

		helper.ToResponseWriter(w, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, e interface{}) {
	webResponse := response.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   nil,
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "application/json")

	helper.ToResponseWriter(w, webResponse)
}
