package dbie

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func makeTestDB(t *testing.T) *bun.DB {
	connStr := fmt.Sprintf("file::memory:%s?cache=shared", t.Name())
	sqldb, err := sql.Open(sqliteshim.DriverName(), connStr)
	if err != nil {
		t.Fatal(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	_, _ = db.Exec("CREATE DATABASE test;")

	_, err = db.NewCreateTable().Model(&user{}).Exec(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return db
}

type user struct {
	Name, LastName, Group string
}

func initUserRepo(t *testing.T) Repo[user] {
	repo := NewRepo[user](
		BunCore[user]{DB: makeTestDB(t), Context: context.Background()},
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
