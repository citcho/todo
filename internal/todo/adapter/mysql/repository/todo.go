package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/citcho/todo/internal/pkg/database"
	"github.com/citcho/todo/internal/todo/adapter/mysql/dao"
	"github.com/citcho/todo/internal/todo/domain/todo"
	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

type TodoRepository struct {
	db *bun.DB
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		db: database.GetDB(),
	}
}

func (tr *TodoRepository) Save(ctx context.Context, t *todo.Todo) error {
	todo := dao.Todo{
		Id:         t.Id(),
		UserId:     t.UserId(),
		Title:      t.Title(),
		Content:    t.Content(),
		IsComplete: t.IsComplete(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
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
	t := dao.Todo{Id: id}
	err := tr.db.NewSelect().
		Model(&t).
		WherePK().
		Scan(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	todo := todo.ReConstructFromRepository(
		t.Id,
		t.UserId,
		t.Title,
		t.Content,
		t.IsComplete,
	)

	return todo, nil
}

func (tr *TodoRepository) Update(ctx context.Context, t *todo.Todo) error {
	todo := dao.Todo{
		Id:         t.Id(),
		UserId:     t.UserId(),
		Title:      t.Title(),
		Content:    t.Content(),
		IsComplete: t.IsComplete(),
	}

	_, err := tr.db.NewUpdate().
		Model(&todo).
		WherePK().
		OmitZero().
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
		Where("user_id = ?", userId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var result []*todo.Todo
	for _, t := range todos {
		result = append(result, todo.ReConstructFromRepository(
			t.Id,
			t.UserId,
			t.Title,
			t.Content,
			t.IsComplete,
		))
	}

	return result, nil
}
