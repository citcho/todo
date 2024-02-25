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
		Completed: t.Completed(),
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

func (tr *TodoRepository) FindById(ctx context.Context, id string) (*todo.Todo, error) {
	var t dao.Todo
	err := tr.db.NewSelect().
		Model(&t).
		Where("ulid = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	todo := todo.ReConstructFromRepository(
		t.Ulid,
		t.UserUlid,
		t.Title,
		t.Content,
		t.Completed,
	)

	return todo, nil
}

func (tr *TodoRepository) Update(ctx context.Context, t *todo.Todo) error {
	todo := dao.Todo{
		Ulid:      t.Ulid(),
		UserUlid:  t.UserId(),
		Title:     t.Title(),
		Content:   t.Content(),
		Completed: t.Completed(),
		UpdatedAt: time.Now(),
	}

	_, err := tr.db.NewUpdate().
		Model(&todo).
		Where("ulid = ?", t.Ulid()).
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

func (tr *TodoRepository) FindAll(ctx context.Context, userId string) ([]*todo.Todo, error) {
	var todos []*dao.Todo
	err := tr.db.NewSelect().
		Model(&todos).
		Where("user_ulid = ?", userId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var result []*todo.Todo
	for _, t := range todos {
		result = append(result, todo.ReConstructFromRepository(
			t.Ulid,
			t.UserUlid,
			t.Title,
			t.Content,
			t.Completed,
		))
	}

	return result, nil
}
