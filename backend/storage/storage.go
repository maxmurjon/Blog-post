package storage

import (
	"context"

	"github.com/maxmurjon/blog_post/models"
)

type StorageRepoI struct {
	User UserRepoI
	// Post PostRepoI
	// Comment CommentRepoI
}

type UserRepoI interface {
	Create(context.Context, models.UserCreateModel) (int, error)
	Update(context.Context, models.UserUpdateModel) (models.User, error)
	GetById(context.Context, models.PrimaryKey) (models.User, error)
	GetList(context.Context, models.PrimaryKey) ([]models.User, error)
	Delete(context.Context, models.PrimaryKey) error
}

type PostRepoI interface {
	Create()
	Update()
	GetById()
	GetList()
	Delete()
}

type CommentRepoI interface {
	Create()
	Update()
	GetById()
	GetList()
	Delete()
}
