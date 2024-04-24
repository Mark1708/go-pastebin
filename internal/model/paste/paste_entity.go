package paste

import "time"

type Paste struct {
	Hash       string    `db:"hash"`
	Title      string    `db:"title"`
	Visibility string    `db:"visibility"`
	CreatedAt  time.Time `db:"created_at"`
	ExpiredAt  time.Time `db:"expired_at"`
	Content    string    `db:"content"`
}

type Pastes []*Paste
