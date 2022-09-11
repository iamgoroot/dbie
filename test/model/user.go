package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       int    `bun:"id,pk,autoincrement" gorm:"primaryKey" pg:",pk" `
	Name     string `bun:"name"`
	LastName string `bun:"last_name" bson:"last_name"`
	Group    string

	gorm.Model `bun:"-" pg:"-" bson:"-"`

	BunCreatedAt *time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" gorm:"-" pg:"-" bson:"-"`
	BunDeletedAt *time.Time `bun:"deleted_at,soft_delete,nullzero" gorm:"-" pg:"-" bson:"-"`

	PgCreatedAt *time.Time `pg:"created_at,notnull,default:current_timestamp" gorm:"-" bun:"-" bson:"-"`
	PgDeletedAt *time.Time `pg:"deleted_at,soft_delete" gorm:"-" bun:"-" bson:"-"`
}
