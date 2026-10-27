package controller

import (
	"net/http"

	"github.com/Benyamin206/backend-clean-arch-go/helper"
	"github.com/Benyamin206/backend-clean-arch-go/model/dto/request"
	"github.com/Benyamin206/backend-clean-arch-go/model/dto/response"
	"github.com/Benyamin206/backend-clean-arch-go/service"
)

type JobController struct {
	JobService service.JobService
}

func NewJobController(jobService service.JobService) *JobController {
	return &JobController{
		JobService: jobService,
	}
}

func (controller *JobController) Add(writer http.ResponseWriter, req *http.Request) {
	jobAddRequest := request.JobAddRequest{}

	helper.DecodeBody(req.Body, &jobAddRequest)

	controller.JobService.Add(req.Context(), jobAddRequest)

	webResponse := response.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   nil,
	}

	helper.ToResponseWriter(writer, webResponse)
}
