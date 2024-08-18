package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/maxmurjon/blog_post/storage"
)

type Store struct {
	db   *sqlx.DB
	user storage.StorageRepoI
}

func NewPostgres(psqlConString string) storage.StorageRepoI {
	db, err := sqlx.Connect("postgres", psqlConString)
	if err != nil {
		log.Panic(err)
	}
	return &Store{
		db: db,
	}
}

func (s *Store) User() storage.UserRepoI{ // postgrsql ga ulangan connection ni userRepo structiga ulanmoqda
	if s.user == nil{
		s.user=&userRepo{db:s.db}
	}
	return s.user
}
