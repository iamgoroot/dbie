package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model   `bun:"-"`
	ID           int    `bun:"id,pk,autoincrement" gorm:"primaryKey"`
	Name         string `bun:"name"`
	LastName     string `bun:"last_name"`
	Group        string
	BunCreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" gorm:"-"`
	BunDeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero" gorm:"-"`
}
