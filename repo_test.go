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

	sqldb, err := sql.Open(sqliteshim.DriverName(), "file::memory:?cache=shared")
	if err != nil {
		t.Fatal(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.Exec("CREATE DATABASE test;")

	type User struct {
		Name, LastName, Group string
	}
	_, err = db.NewCreateTable().Model(&User{}).Exec(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewRepo[User](BunBackend[User]{DB: db, Context: context.Background()})

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
	//Select Page
	usersFromGroups1And6, err := repo.SelectPage(Page{
		Limit:  3,
		Offset: 1,
	}, `"user"."group"`, In, []string{"group1", "group6"}, Sort{
		Field: "last_name",
		Order: ASC,
	}, Sort{
		Field: "name",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(usersFromGroups1And6)

	// Select One
	one, err := repo.SelectOne(`"user"."last_name"`, Eq, "User6LastName")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(one)

	// Select that doesn't exist
	one, err = repo.SelectOne(`"user"."last_name"`, Eq, "EXPECT ERROR BECAUSE I DONT EXIST")
	if err == nil {
		t.Fatal("expected error found user", one)
	}
	usersFromGroup2, err := repo.Select(`"user"."group"`, Eq, "group2")
	if err != nil {
		t.Fatal(err)
	}
	if len(usersFromGroup2) != 2 {
		t.Fatal("expected two users in group2 but found", usersFromGroup2)
	}
	//TODO: assert
}
