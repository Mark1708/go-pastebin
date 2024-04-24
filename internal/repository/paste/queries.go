package paste

import "github.com/Mark1708/go-pastebin/pkg/db"

var (
	getPasteByHashQuery = db.Query{
		Name:     "paste_repository.GetByHash",
		QueryRaw: "SELECT * FROM paste WHERE hash=$1 LIMIT 1;",
	}
	createPasteQuery = db.Query{
		Name: "paste_repository.Create",
		QueryRaw: `
			INSERT INTO paste (
				hash, title, visibility,
				created_at, expired_at, content
			)
			VALUES (
				$1, $2, $3,
				$4, $5, $6
			)
			RETURNING
				hash, title, visibility,
				created_at, expired_at, content;
		`,
	}
	updatePasteByHashQuery = db.Query{
		Name: "paste_repository.UpdateByHash",
		QueryRaw: `
			UPDATE paste
			SET
				title = COALESCE($1, title),
				visibility = COALESCE($2, visibility),
				expired_at = COALESCE($3, expired_at),
				content = COALESCE($4, content)
			WHERE hash = $5
			RETURNING
				hash, title, visibility,
				created_at, expired_at, content;
		`,
	}
	deletePasteByHashQuery = db.Query{
		Name:     "paste_repository.DeleteByHash",
		QueryRaw: "DELETE FROM paste WHERE hash = $1;",
	}
)
