package test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/iamgoroot/dbie"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"testing"
)

func TestNewUserRepo(t *testing.T) {

	//repo := dbie.NewRepo[user](
	//	dbie.BunCore[user]{DB: makeTestDB(t), Context: context.Background()},
	//)
	//err := repo.Insert(createUsers()...)
	//if err != nil {
	//	t.Fatal(err)
	//}
	////Select Page
	//usersFromGroups1And6, err := repo.SelectPage(dbie.Page{
	//	Limit:  3,
	//	Offset: 1,
	//}, `"user"."group"`, dbie.In, []string{"group1", "group6"}, dbie.Sort{
	//	Field: "last_name",
	//	Order: dbie.ASC,
	//}, dbie.Sort{
	//	Field: "name",
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(usersFromGroups1And6)
	//
	//// Select One
	//one, err := repo.SelectOne(`"user"."last_name"`, dbie.Eq, "UserLastName6")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(one)
	//
	//// Select that doesn't exist
	//one, err = repo.SelectOne(`"user"."last_name"`, dbie.Eq, "EXPECT ERROR BECAUSE I DONT EXIST")
	//if err == nil {
	//	t.Fatal("expected error found user", one)
	//}
	//usersFromGroup2, err := repo.Select(`"user"."group"`, dbie.Eq, "group2")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if len(usersFromGroup2) != 2 {
	//	t.Fatal("expected two users in group2 but found", usersFromGroup2)
	//}
	//TODO: assert
}

func makeTestDB(t *testing.T) *bun.DB {
	sqldb, err := sql.Open(sqliteshim.DriverName(), fmt.Sprintf("file::memory:?cache=shared"))
	if err != nil {
		t.Fatal(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.Exec("CREATE DATABASE test;")

	_, err = db.NewCreateTable().Model(&user{}).Exec(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return db
}

type user struct {
	Name, LastName, Group string
}

func initUserRepo(t *testing.T) dbie.Repo[user] {
	repo := dbie.NewRepo[user](
		dbie.BunCore[user]{DB: makeTestDB(t), Context: context.Background()},
	)
	err := repo.Insert(createUsers()...)
	if err != nil {
		t.Fatal(err)
	}
	return repo
}

func createUsers() (users []user) {
	const n = 100
	users = make([]user, n)
	for i := 0; i < n; i++ {
		users[i] = user{
			Name:     fmt.Sprintf("User%dName", i),
			LastName: fmt.Sprintf("UserLastName%d", i),
			Group:    fmt.Sprintf("group%d", i%10),
		}
	}
	return users
}

func assertUser(t *testing.T, expected user, got user) {
	switch {
	case got.Group != expected.Group:
		t.Fatal("expected singleUser.Group", expected.Group, "got", got.Group)
	case got.Name != expected.Name:
		t.Fatal("expected singleUser.Name", expected.Name, "got", got.Name)
	case got.LastName != expected.LastName:
		t.Fatal("expected singleUser.LastName", expected.LastName, "got", got.LastName)
	}
}
