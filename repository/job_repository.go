package repository

import (
	"context"
	"database/sql"

	"github.com/Benyamin206/backend-clean-arch-go/model/domain"
)

type JobRepository interface {
	Add(ctx context.Context, tx *sql.Tx, job domain.Job)
}
