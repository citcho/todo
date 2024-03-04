package dao

import (
	"time"

	"github.com/hexisa_go_nal_todo/internal/user/adapter/mysql/dao"
)

type Todo struct {
	Id         string    `bun:",pk,type:char(26)"`
	User       *dao.User `bun:",rel:belongs-to"`
	UserId     string    `bun:",type:char(26),notnull"`
	Title      string    `bun:",notnull"`
	Content    string    `bun:",type:text,notnull"`
	IsComplete bool      `bun:",type:int,notnull"`
	CreatedAt  time.Time `bun:",notnull"`
	UpdatedAt  time.Time `bun:",notnull"`
	DeletedAt  time.Time `bun:",soft_delete,nullzero"`
}
