package dbie

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSelectOne(t *testing.T) {
	testWithRepos(t, func(t *testing.T, repo Repo[testUser]) {
		// Select One
		singleUser, err := repo.SelectOne("last_name", Eq, "UserLastName16")
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

func getCoreName(repository Repo[testUser]) string {
	repo := GenericBackend[testUser]{repository}
	switch core := repo.Core.(type) {
	case BunCore[testUser]:
		return "bun"
	case GormCore[testUser]:
		return "gorm"
	case GenericBackend[testUser]:
		if core.Core != nil {
			return getCoreName(core.Core)
		}
		return reflect.TypeOf(core).String()
	default:
		return reflect.TypeOf(core).String()
	}
}
func testWithRepos(t *testing.T, testFunc func(*testing.T, Repo[testUser])) {
	repos := []Repo[testUser]{initBunRepo(t), initGormRepo(t)}
	for _, r := range repos {
		err := r.Insert(createUsers()...)
		if err != nil {
			t.Fatal(err)
		}
		repo := r // capture var
		name := fmt.Sprintf("db[%s]", getCoreName(repo))
		t.Run(name, func(t *testing.T) {
			defer repo.Close()
			t.Parallel()
			testFunc(t, repo)
		})
	}
}
func TestSelectOneNoRows(t *testing.T) {
	testWithRepos(t, func(t *testing.T, repo Repo[testUser]) {
		// Select that doesn't exist
		notFoundUser, err := repo.SelectOne("last_name", Eq, "EXPECT ERROR BECAUSE I DONT EXIST")

		// Expect dbie.ErrNoRows
		if !errors.Is(err, ErrNoRows) {
			t.Fatal("expected ErrNoRows", notFoundUser, "error", err)
		}
		// Expect empty model
		if !reflect.DeepEqual(notFoundUser, testUser{}) {
			t.Fatal("expected empty testUser but found", notFoundUser, "error", err)
		}
	})
}

func TestGormSelectOne(t *testing.T) {
	testWithRepos(t, func(t *testing.T, repo Repo[testUser]) {
		// Select One
		singleUser, err := repo.SelectOne("last_name", Eq, "UserLastName16")
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
	testWithRepos(t, func(t *testing.T, repo Repo[testUser]) {
		// Select that doesn't exist
		notFoundUser, err := repo.SelectOne("last_name", Eq, "EXPECT ERROR BECAUSE I DONT EXIST")

		// Expect dbie.ErrNoRows
		if !errors.Is(err, ErrNoRows) {
			t.Fatal("expected ErrNoRows", notFoundUser, "error", err)
		}
		// Expect empty model
		if !reflect.DeepEqual(notFoundUser, testUser{}) {
			t.Fatal("expected empty testUser but found", notFoundUser, "error", err)
		}
	})
}
