package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrAlreadyExists = errors.New("already exists")
)

func ConvertPGError(err error) error {
	if err == nil {
		return nil
	}

	// https://github.com/jackc/pgx/wiki/Error-Handling
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		logger.Error(context.TODO(), pgErr.Message) // => syntax error at end of input
		logger.Error(context.TODO(), pgErr.Code)    // => 42601

		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return fmt.Errorf("%s: %w", pgErr.Message, ErrAlreadyExists)
		default:
			return err
		}
	}
	return err
}
