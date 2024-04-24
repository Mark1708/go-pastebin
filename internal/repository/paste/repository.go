package paste

import (
	"context"
	"time"

	"github.com/Mark1708/go-pastebin/internal/model/paste"
	"github.com/Mark1708/go-pastebin/pkg/db"
)

type Repository interface {
	GetPasteByHash(ctx context.Context, hash string) (paste.Paste, error)

	CreatePaste(ctx context.Context, entity paste.Paste) (paste.Paste, error)

	UpdatePaste(ctx context.Context, entity paste.Paste) (paste.Paste, error)

	DeletePaste(ctx context.Context, hash string) error
}

type repository struct {
	db db.Client
	tx db.TxManager
}

func NewRepository(db db.Client, tx db.TxManager) Repository {
	return &repository{db: db, tx: tx}
}

func (r *repository) GetPasteByHash(ctx context.Context, hash string) (paste.Paste, error) {
	entity := paste.Paste{}
	err := r.db.DB().QueryRowContext(ctx, getPasteByHashQuery, hash).Scan(
		&entity.Hash,
		&entity.Title,
		&entity.Visibility,
		&entity.CreatedAt,
		&entity.ExpiredAt,
		&entity.Content,
	)
	if err != nil {
		return paste.Paste{}, err
	}

	return entity, err
}

func (r *repository) CreatePaste(ctx context.Context, entity paste.Paste) (paste.Paste, error) {
	err := r.tx.Serializable(
		ctx,
		func(ctx context.Context) error {
			err := r.db.DB().QueryRowContext(
				ctx, createPasteQuery,
				entity.Hash,
				entity.Title,
				entity.Visibility,
				entity.CreatedAt.Format(time.RFC3339),
				entity.ExpiredAt.Format(time.RFC3339),
				entity.Content,
			).Scan(
				&entity.Hash,
				&entity.Title,
				&entity.Visibility,
				&entity.CreatedAt,
				&entity.ExpiredAt,
				&entity.Content,
			)
			return err
		},
	)
	if err != nil {
		return paste.Paste{}, err
	}
	return entity, nil
}

func (r *repository) UpdatePaste(ctx context.Context, entity paste.Paste) (paste.Paste, error) {
	err := r.tx.Serializable(
		ctx,
		func(ctx context.Context) error {
			err := r.db.DB().QueryRowContext(
				ctx, updatePasteByHashQuery,
				entity.Title,
				entity.Visibility,
				entity.CreatedAt.Format(time.RFC3339),
				entity.Content,
				entity.Hash,
			).Scan(
				&entity.Hash,
				&entity.Title,
				&entity.Visibility,
				&entity.CreatedAt,
				&entity.ExpiredAt,
				&entity.Content,
			)
			return err
		},
	)
	if err != nil {
		return paste.Paste{}, err
	}
	return entity, nil
}

func (r *repository) DeletePaste(ctx context.Context, hash string) error {
	return r.tx.Serializable(
		ctx,
		func(ctx context.Context) error {
			commandTag, err := r.db.DB().ExecContext(
				ctx, deletePasteByHashQuery,
				hash,
			)
			if err != nil {
				return err
			}
			_ = commandTag.Delete()
			return err
		},
	)
}
