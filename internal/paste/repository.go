package paste

import "github.com/Mark1708/go-pastebin/internal/paste/models"

type Repository interface {
	GetByHash(hash string) (models.Paste, error)

	Create(paste models.Paste) (models.Paste, error)

	Update(paste models.Paste) (models.Paste, error)

	Delete(hash string) error
}
