package pg

import (
	"context"

	"github.com/Mark1708/go-pastebin/pkg/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type key string

const (
	TxKey key = "tx"
)

var ErrNotFound = errors.New("error not found")
var ErrScan = errors.New("error scanning")
var ErrQuery = errors.New("error db query")
var ErrExec = errors.New("error db execution")

type pg struct {
	dbc *pgxpool.Pool
}

func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{
		dbc: dbc,
	}
}

func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return ErrQuery
	}
	err = row.Scan(dest)
	if err != nil {
		return ErrScan
	}
	return nil
}

func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return ErrQuery
	}
	err = rows.Scan(dest)
	if err != nil {
		return ErrScan
	}
	return nil
}

func (p *pg) ExecContext(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	var commandTag pgconn.CommandTag
	var err error

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		commandTag, err = tx.Exec(ctx, q.QueryRaw, args...)
	} else {
		commandTag, err = p.dbc.Exec(ctx, q.QueryRaw, args...)
	}

	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return pgconn.CommandTag{}, err
		}
		return pgconn.CommandTag{}, ErrExec
	}
	return commandTag, nil
}

func (p *pg) QueryContext(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	var rows pgx.Rows
	var err error

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		rows, err = tx.Query(ctx, q.QueryRaw, args...)
	} else {
		rows, err = p.dbc.Query(ctx, q.QueryRaw, args...)
	}

	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return nil, err
	}

	return rows, err
}

func (p *pg) QueryRowContext(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p *pg) Close() {
	p.dbc.Close()
}

func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}
