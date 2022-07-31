package test

import (
	"context"
	"database/sql"
	"github.com/iamgoroot/dbie/core/test/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	mysqlGorm "gorm.io/driver/mysql"
	pgGorm "gorm.io/driver/postgres"
	sqliteGorm "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dsn = []string{
	"postgres://postgres:postgres@localhost:5432/test ",
	"user:pass@/test",
	"sqlserver://sa:passWORD1@localhost:14339?database=test",
	"file::memory:?cache=shared",
}

func makeBunSqlite(dsn string) *bun.DB {
	sqldb, _ := sql.Open(sqliteshim.DriverName(), dsn)
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.Exec("CREATE DATABASE repo;")
	db.NewDelete().Model(&model.User{}).Exec(context.Background())
	_, _ = db.NewCreateTable().Model(&model.User{}).Exec(context.Background())
	return db
}

func makeGormSqlite(dsn string) *gorm.DB {
	db, _ := gorm.Open(sqliteGorm.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	db.Where("1 = 1").Delete(&model.User{})
	return db
}
func makeGormMysql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysqlGorm.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	db.Where("1 = 1").Delete(&model.User{})
	return db
}
func makeBunPostgres(dsn string) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(true),
	))
	db.NewDelete().Model(&model.User{}).Exec(context.Background())
	_, _ = db.NewCreateTable().Model(&model.User{}).Exec(context.Background())
	return db
}
func makeGormPostgres(dsn string) *gorm.DB {
	db, _ := gorm.Open(pgGorm.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	db.Where("1 = 1").Delete(&model.User{})
	return db
}

func makeBunMysql(dsn string) *bun.DB {
	sqldb, _ := sql.Open("mysql", dsn)
	db := bun.NewDB(sqldb, mysqldialect.New())
	db.NewDelete().Model(&model.User{}).Exec(context.Background())
	_, _ = db.NewCreateTable().Model(&model.User{}).Exec(context.Background())
	return db
}
