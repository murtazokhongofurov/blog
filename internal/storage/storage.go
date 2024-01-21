package storage

import (
	"database/sql"

	Db "github.com/blog/internal/storage/db"
	"github.com/blog/internal/storage/repo"
)

type StorageI interface {
	Blog() repo.BlogService
}

type StoragePg struct {
	db       *sql.DB
	blogRepo repo.BlogService
}

func NewStoragePg(db *sql.DB) *StoragePg {
	return &StoragePg{
		db:       db,
		blogRepo: Db.NewStorage(db),
	}
}

func (s StoragePg) Blog() repo.BlogService {
	return s.blogRepo
}
