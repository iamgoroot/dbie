package dbie

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"testing"
)

func TestNewUserRepo(t *testing.T) {
	type User struct {
		Name, LastName, Group string
	}
	type UserRepo struct {
		Repo[User]
	}
	sqldb, err := sql.Open(sqliteshim.DriverName(), "file::memory:?cache=shared")
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.Exec("CREATE DATABASE test;")

	_, err = db.NewCreateTable().Model(&User{}).Exec(context.Background())
	fmt.Println("table create err", err)
	repo := UserRepo{
		BunBackend[User]{
			Context: context.Background(),
			DB:      db,
		},
	}
	err = repo.Insert(User{
		Name:     "User1Name",
		LastName: "User1LastName",
		Group:    "group1",
	}, User{
		Name:     "User2Name",
		LastName: "User2LastName",
		Group:    "group1",
	}, User{
		Name:     "User3Name",
		LastName: "User3LastName",
		Group:    "group1",
	}, User{
		Name:     "User4Name",
		LastName: "User4LastName",
		Group:    "group2",
	}, User{
		Name:     "User5Name",
		LastName: "User5LastName",
		Group:    "group2",
	}, User{
		Name:     "User6Name",
		LastName: "User6LastName",
		Group:    "group6",
	})
	if err != nil {
		t.Fatal(err)
	}
	usersFromGroups1And6, err := repo.SelectPage(Page{
		Limit:  3,
		Offset: 1,
	}, `"user"."group"`, In, []string{"group1", "group6"}, Sort{
		Field: "last_name",
		Desc:  true,
	}, Sort{
		Field: "name",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(usersFromGroups1And6)
}
