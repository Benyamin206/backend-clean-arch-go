package repository

import (
	"context"
	"database/sql"

	"github.com/Benyamin206/backend-clean-arch-go/helper"
	"github.com/Benyamin206/backend-clean-arch-go/model/domain"
)

type JobRepositoryImpl struct {
}

func NewJobRepository() JobRepository {
	return &JobRepositoryImpl{}
}

func (repository *JobRepositoryImpl) Add(ctx context.Context, tx *sql.Tx, job domain.Job) {
	sql := "insert into jobs values ?"

	_, err := tx.ExecContext(ctx, sql, job.Name)

	helper.PanicIfError(err)

}
