package dao

import (
	"time"

	"github.com/hexisa_go_nal_todo/internal/user/adapter/mysql/dao"
)

type Todo struct {
	Ulid      string    `bun:",pk,type:char(26)"`
	User      *dao.User `bun:",rel:belongs-to"`
	UserUlid  string    `bun:",pk,type:char(26),notnull"`
	Title     string    `bun:",notnull"`
	Content   string    `bun:",type:text,notnull"`
	Status    string    `bun:",type:varchar(255),notnull"`
	CreatedAt time.Time `bun:",notnull"`
	UpdatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}
