package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hexisa_go_nal_todo/internal/pkg/database"
	"github.com/hexisa_go_nal_todo/internal/user/adapter/mysql/dao"
	"github.com/hexisa_go_nal_todo/internal/user/domain/user"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDB(),
	}
}

func (ur *UserRepository) Exists(ctx context.Context, u *user.User) (bool, error) {
	exists, err := ur.db.NewSelect().
		Model((*dao.User)(nil)).
		Where("email = ?", u.Email()).
		Exists(ctx)
	if err != nil {
		panic(err)
	}

	return exists, nil
}

func (ur *UserRepository) Save(ctx context.Context, u *user.User) error {
	user := dao.User{
		Id:        u.Id(),
		Name:      u.Name(),
		Email:     u.Email(),
		Password:  u.Password(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := ur.db.NewInsert().
		Model(&user).
		Exec(ctx)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			log.Println(err)
			return errors.New("予期せぬエラーが発生しました。管理者にお問い合わせください。")
		}

		return err
	}

	return nil
}

func (ur *UserRepository) FetchByEmail(ctx context.Context, email string) (*user.User, error) {
	var u dao.User
	err := ur.db.NewSelect().
		Model(&u).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return &user.User{}, err
	}

	domainUser := user.ReConstructFromRepository(
		u.Id,
		u.Name,
		u.Email,
		u.Password,
	)

	return domainUser, nil
}

func (ur *UserRepository) FetchById(ctx context.Context, id string) (*user.User, error) {
	var u dao.User
	err := ur.db.NewSelect().
		Model(&u).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return &user.User{}, err
	}

	domainUser := user.ReConstructFromRepository(
		u.Id,
		u.Name,
		u.Email,
		u.Password,
	)

	return domainUser, nil
}
