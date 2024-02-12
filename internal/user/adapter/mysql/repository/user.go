package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hexisa_go_nal_todo/internal/user/adapter/mysql/dao"
	"github.com/hexisa_go_nal_todo/internal/user/domain/user"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{
		db: db,
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
		Ulid:      u.Ulid(),
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
