package app

import (
	"github.com/Benyamin206/backend-clean-arch-go/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.PanicHandler = exception.ErrorHandler

	return router
}
