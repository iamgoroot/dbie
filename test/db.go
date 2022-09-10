package test

import (
	"context"
	"database/sql"
	"github.com/go-pg/pg/extra/pgdebug/v10"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/iamgoroot/dbie/core/test/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	pgGorm "gorm.io/driver/postgres"
	sqliteGorm "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

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

func makePg(dsn string) *pg.DB {
	opt, err := pg.ParseURL(dsn)
	if err != nil {
		log.Fatalln(err)
	}
	db := pg.Connect(opt)
	db.AddQueryHook(pgdebug.NewDebugHook())
	db.Model(&model.User{}).Context(context.Background()).DropTable(&orm.DropTableOptions{Cascade: true})
	err = db.Model(&model.User{}).Context(context.Background()).CreateTable(&orm.CreateTableOptions{FKConstraints: true})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
func makeGormPostgres(dsn string) *gorm.DB {
	db, _ := gorm.Open(pgGorm.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	db.Where("1 = 1").Delete(&model.User{})
	return db
}

// TODO: fix Mysql
//func makeBunMysql(dsn string) *bun.DB {
//	sqldb, _ := sql.Open("mysql", dsn)
//	db := bun.NewDB(sqldb, mysqldialect.New())
//	db.NewDelete().Model(&model.User{}).Exec(context.Background())
//	_, _ = db.NewCreateTable().Model(&model.User{}).Exec(context.Background())
//	return db
//}

// TODO: fix Mysql
//func makeGormMysql(dsn string) *gorm.DB {
//	db, Err := gorm.Open(mysqlGorm.Open(dsn), &gorm.Config{})
//	if Err != nil {
//		panic(Err)
//	}
//	db.AutoMigrate(&model.User{})
//	db.Where("1 = 1").Delete(&model.User{})
//	return db
//}
