package test

import (
	"context"
	"errors"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core/test/model"
	"github.com/iamgoroot/dbie/core/test/repo"
	"reflect"
	"testing"
)

func TestSelectOne(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo dbie.Repo[model.User]) {
		// Select One
		singleUser, err := repo.SelectOne("last_name", dbie.Eq, "UserLastName16")
		if err != nil {
			t.Fatal(err, repo)
		}
		// Assert
		expectedGroup := "group6"
		expectedName := "User16Name"
		switch {
		case singleUser.Group != expectedGroup:
			t.Fatal("expected singleUser.Group", expectedGroup, "got", singleUser.Group)
		case singleUser.Name != expectedName:
			t.Fatal("expected singleUser.Name", expectedName, "got", singleUser.Name)
		}
	})
}

func testAllCores(t *testing.T, testFunc func(*testing.T, dbie.Repo[model.User])) {
	makers := map[string]func(ctx context.Context) repo.User{
		"BunSqlite":    repo.Bun{DB: makeBunSqlite("file::memory:?")}.NewBunUser,
		"GormSqlite":   repo.Gorm{DB: makeGormSqlite("file::memory:?")}.NewGormUser,
		"BunPostgres":  repo.Bun{DB: makeBunPostgres("postgres://user:pass@127.0.0.1:5432/test?sslmode=disable")}.NewBunUser,
		"GormPostgres": repo.Gorm{DB: makeGormPostgres("postgres://user:pass@127.0.0.1:5433/test?sslmode=disable")}.NewGormUser,
		//repo.Gorm{DB: makeGormMysql("user:pass@tcp(localhost:3307)/test")},
		//repo.Bun{DB: makeBunMysql("user:pass@tcp(localhost:3306)/test")},

	}
	for key, maker := range makers {
		repo := maker(context.Background())
		repo.Init()
		err := repo.Insert(createUsers()...)
		if err != nil {
			t.Fatal(err)
		}
		defer repo.Close()
		t.Run(key, func(t *testing.T) {
			//t.Parallel()
			testFunc(t, repo)
		})
	}
}

func TestSelectOneNoRows(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo dbie.Repo[model.User]) {
		// Select that doesn't exist
		notFoundUser, err := repo.SelectOne("last_name", dbie.Eq, "EXPECT ERROR BECAUSE I DONT EXIST")

		// Expect dbie.ErrNoRows
		if !errors.Is(err, dbie.ErrNoRows) {
			t.Fatal("expected ErrNoRows", notFoundUser, "error", err)
		}
		// Expect empty model
		if !reflect.DeepEqual(notFoundUser, model.User{}) {
			t.Fatal("expected empty model.User but found", notFoundUser, "error", err)
		}
	})
}

func TestGormSelectOne(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo dbie.Repo[model.User]) {
		// Select One
		singleUser, err := repo.SelectOne("last_name", dbie.Eq, "UserLastName16")
		if err != nil {
			t.Fatal(err)
		}

		// Assert
		expectedGroup := "group6"
		expectedName := "User16Name"
		switch {
		case singleUser.Group != expectedGroup:
			t.Fatal("expected singleUser.Group", expectedGroup, "got", singleUser.Group)
		case singleUser.Name != expectedName:
			t.Fatal("expected singleUser.Name", expectedName, "got", singleUser.Name)
		}
	})
}

func TestGormSelectOneNoRows(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo dbie.Repo[model.User]) {
		// Select that doesn't exist
		notFoundUser, err := repo.SelectOne("last_name", dbie.Eq, "EXPECT ERROR BECAUSE I DONT EXIST")

		// Expect dbie.ErrNoRows
		if !errors.Is(err, dbie.ErrNoRows) {
			t.Fatal("expected ErrNoRows", notFoundUser, "error", err)
		}
		// Expect empty model
		if !reflect.DeepEqual(notFoundUser, model.User{}) {
			t.Fatal("expected empty model.User but found", notFoundUser, "error", err)
		}
	})
}
