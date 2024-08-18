package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/maxmurjon/blog_post/models"
)

type userRepo struct {
	db *sqlx.DB
}

func (u *userRepo) Create(ctx context.Context, user models.UserCreateModel) (id int, err error) {
	return id, err
}
func (u *userRepo) Update(context.Context, models.UserUpdateModel) (user models.User, err error) {
	return user, err
}

func (u userRepo) GetById(context.Context, models.PrimaryKey) (user models.User, err error) {
	return user, err
}

func (u userRepo) GetList(context.Context, models.PrimaryKey) (users []models.User, err error) {
	return users, err
}

func (u userRepo) Delete(context.Context, models.PrimaryKey) (err error) {
	return err
}
