package service

import (
	"context"

	"github.com/Benyamin206/backend-clean-arch-go/model/dto/request"
)

type JobService interface {
	Add(ctx context.Context, req request.JobAddRequest)
}
