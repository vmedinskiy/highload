package dbutils

import (
	"context"
	"database/sql"

	"github.com/go-faster/errors"
	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"go.uber.org/multierr"
)

type TxFunc func(tx *sqlx.Tx) error

type DB interface {
	BeginTxx(context.Context, *sql.TxOptions) (*sqlx.Tx, error)
}

func sqlerr(err error, query string) error {
	if err == nil {
		return nil
	}
	return errors.Wrapf(err, "query: %s", query)
}

func Named(q string, arg any) (nq string, args []any, err error) {
	nq, args, err = sqlx.Named(q, arg)
	args = append(
		[]any{pgx.QuerySimpleProtocol(true)},
		args...,
	)
	return nq, args, sqlerr(err, q)
}

func NamedExec(ctx context.Context, db sqlx.ExtContext, q string, arg any) error {
	nq, args, err := Named(q, arg)
	if err != nil {
		return err
	}

	nq = db.Rebind(nq)
	_, err = db.ExecContext(ctx, nq, args...)

	return sqlerr(err, nq)
}

func NamedGet(ctx context.Context, db sqlx.ExtContext, dest any, q string, arg any) error {
	nq, args, err := db.BindNamed(q, arg)
	if err != nil {
		return sqlerr(err, q)
	}

	args = append(
		[]any{pgx.QuerySimpleProtocol(true)},
		args...,
	)
	err = sqlx.GetContext(ctx, db, dest, nq, args...)
	return sqlerr(err, nq)
}

func NamedSelect(ctx context.Context, db sqlx.ExtContext, dest any, q string, arg any) error {
	nq, args, err := Named(q, arg)
	if err != nil {
		return err
	}

	nq = db.Rebind(nq)
	err = sqlx.SelectContext(ctx, db, dest, nq, args...)

	return sqlerr(err, nq)
}

func WrapTxx(ctx context.Context, db DB, opts *sql.TxOptions, f TxFunc) (err error) {
	var tx *sqlx.Tx

	tx, err = db.BeginTxx(ctx, opts)
	if err != nil {
		return errors.Errorf("%w, begin transaction", err)
	}
	defer func() {
		r := recover()
		if r != nil {
			_ = tx.Rollback()
			panic(r)
		}

		if err != nil {
			err = multierr.Combine(err, tx.Rollback())
		} else {
			err = tx.Commit()
		}
	}()

	return f(tx)
}
