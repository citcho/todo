package dao

import "time"

type User struct {
	Ulid      string    `bun:",pk,type:char(26)"`
	Name      string    `bun:",notnull"`
	Email     string    `bun:",notnull,unique"`
	Password  string    `bun:",notnull"`
	CreatedAt time.Time `bun:",notnull"`
	UpdatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}
