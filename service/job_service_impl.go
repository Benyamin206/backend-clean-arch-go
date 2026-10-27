package service

import (
	"context"
	"database/sql"

	"github.com/Benyamin206/backend-clean-arch-go/helper"
	"github.com/Benyamin206/backend-clean-arch-go/model/domain"
	"github.com/Benyamin206/backend-clean-arch-go/model/dto/request"
	"github.com/Benyamin206/backend-clean-arch-go/repository"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
)

type JobServiceImpl struct {
	JobRepository repository.JobRepository
	DB            *sql.DB
	Validator     *validator.Validate
}

func NewJobService(jobRepository repository.JobRepository, db *sql.DB, validator *validator.Validate) JobService {
	return &JobServiceImpl{
		JobRepository: jobRepository,
		DB:            db,
		Validator:     validator,
	}
}

func (service *JobServiceImpl) Add(ctx context.Context, req request.JobAddRequest) {
	err := service.Validator.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()

	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	job := domain.Job{
		Name: req.Name,
	}

	service.JobRepository.Add(ctx, tx, job)
}
