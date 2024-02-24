package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hexisa_go_nal_todo/internal/todo/adapter/mysql/dao"
	"github.com/hexisa_go_nal_todo/internal/todo/domain/todo"
	"github.com/uptrace/bun"
)

type TodoRepository struct {
	db *bun.DB
}

func NewTodoRepository(db *bun.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (tr *TodoRepository) Save(ctx context.Context, t *todo.Todo) error {
	todo := dao.Todo{
		Ulid:      t.Ulid(),
		UserUlid:  t.UserId(),
		Title:     t.Title(),
		Content:   t.Content(),
		Status:    t.Status().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := tr.db.NewInsert().
		Model(&todo).
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
