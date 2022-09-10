package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       int    `bun:"id,pk,autoincrement" gorm:"primaryKey" pg:",pk"`
	Name     string `bun:"name"`
	LastName string `bun:"last_name"`
	Group    string

	gorm.Model `bun:"-" pg:"-"`

	BunCreatedAt *time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" gorm:"-" pg:"-"`
	BunDeletedAt *time.Time `bun:"deleted_at,soft_delete,nullzero" gorm:"-" pg:"-"`

	PgCreatedAt *time.Time `pg:"created_at,notnull,default:current_timestamp" gorm:"-" bun:"-"`
	PgDeletedAt *time.Time `pg:"deleted_at,soft_delete" gorm:"-" bun:"-"`
}
