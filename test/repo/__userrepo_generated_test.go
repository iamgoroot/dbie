package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core/test/model"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"

	"testing"
)

func TestNewUserRepo(t *testing.T) {
	for _, v := range []interface {
		NewUser(ctx context.Context) User
	}{
		Bun{DB: makeBunSqlite()},
		Gorm{DB: makeGormSqltie()},
	} {
		repo := v.NewUser(context.Background())
		err := repo.Insert(
			model.User{
				Name:     "User1Name",
				LastName: "User1LastName",
				Group:    "group1",
			}, model.User{
				Name:     "User2Name",
				LastName: "User2LastName",
				Group:    "group1",
			}, model.User{
				Name:     "User3Name",
				LastName: "User3LastName",
				Group:    "group1",
			}, model.User{
				Name:     "User4Name",
				LastName: "User4LastName",
				Group:    "group2",
			}, model.User{
				Name:     "User5Name",
				LastName: "User5LastName",
				Group:    "group2",
			}, model.User{
				Name:     "User6Name",
				LastName: "User6LastName",
				Group:    "group6",
			})
		if err != nil {
			t.Fatal(err)
		}
		usersFromGroups1And6, err := repo.SelectByGroupIn(dbie.Page{
			Limit: 20,
		}, "group1", "group6")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(usersFromGroups1And6)

		singleUser, err := repo.SelectByGroupOrderByNameDescOrderByIDAsc("group1")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(singleUser)
	}

}

func makeBunSqlite() *bun.DB {
	sqldb, _ := sql.Open(sqliteshim.DriverName(), "file::memory:?cache=shared")
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.Exec("CREATE DATABASE repo;")
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	_, _ = db.NewCreateTable().Model(&model.User{}).Exec(context.Background())
	return db
}

func makeGormSqltie() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})
	return db
}
