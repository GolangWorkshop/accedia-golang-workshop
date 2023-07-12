package store

import (
	"database/sql"

	"github.com/GolangWorkshop/library/util"
	baseLog "github.com/rs/zerolog/log"
)

var log = baseLog.With().Str("package", "store").Logger()

type store struct {
	db *sql.DB
}

type Store interface {
	CreateBook(*Book) (*Book, error)
	GetBooks() (*[]Book, error)
	GetBookById(id string) (*Book, error)
	UpdateBookById(id string, book *Book) (*Book, error)
	DeleteBookById(id string) (int64, error)
	Register(*User) (*User, error)
	Login(*User) (*util.JwtInfo, error)
}

func NewStore(db *sql.DB) Store {
	return &store{
		db: db,
	}
}
