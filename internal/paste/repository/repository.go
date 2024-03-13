package repository

import (
	"github.com/Mark1708/go-pastebin/internal/paste/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
}

func (r Repository) GetByHash(hash string) (models.Paste, error) {
	paste := models.Paste{}
	err := r.DB.Get(&paste, "SELECT * FROM paste WHERE hash=$1", hash)
	return paste, err
}

func (r Repository) Create(paste models.Paste) (models.Paste, error) {
	tx := r.DB.MustBegin()

	_, insertErr := tx.NamedExec(
		`
		INSERT INTO paste 
		    (hash, title, visibility, created_at, expired_at, content) 
		VALUES 
		    (:hash, :title, :visibility, :created_at, :expired_at, :content)
		`,
		&paste,
	)
	if insertErr != nil {
		return models.Paste{}, insertErr
	}

	txErr := tx.Commit()
	if txErr != nil {
		return models.Paste{}, txErr
	}
	return paste, nil
}

func (r Repository) Update(paste models.Paste) (models.Paste, error) {
	tx := r.DB.MustBegin()

	_, insertErr := tx.NamedExec(
		`
		UPDATE paste 
		SET 
		 title = :title,
		 visibility = :visibility,
		 created_at = :created_at,
		 expired_at = :expired_at,
		 content = :content
		WHERE hash = :hash
		`,
		&paste,
	)
	if insertErr != nil {
		return models.Paste{}, insertErr
	}

	txErr := tx.Commit()
	if txErr != nil {
		return models.Paste{}, txErr
	}
	return paste, nil
}

func (r Repository) Delete(hash string) error {
	tx := r.DB.MustBegin()

	tx.MustExec("DELETE FROM paste WHERE hash = $1", hash)

	txErr := tx.Commit()
	if txErr != nil {
		return txErr
	}
	return nil
}
