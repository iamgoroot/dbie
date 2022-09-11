package test

import (
	"errors"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core/test/model"
	"github.com/iamgoroot/dbie/core/test/repo"
	"reflect"
	"testing"
)

func TestSelectOne(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo repo.User) {
		// Select One
		singleUser, err := repo.SelectOne("last_name", dbie.Eq, "UserLastName16")
		if err != nil {
			t.Fatal(err, repo)
		}
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

func TestSelectOneNoRows(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo repo.User) {
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
	testAllCores(t, func(t *testing.T, repo repo.User) {
		// Select One
		singleUser, err := repo.SelectOne("last_name", dbie.Eq, "UserLastName16")
		if err != nil {
			t.Fatal(err)
		}
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
	testAllCores(t, func(t *testing.T, repo repo.User) {
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

func TestSelect(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo repo.User) {
		// Select slice
		users, err := repo.Select(`last_name`, dbie.In, []string{"UserLastName16", "UserLastName17"})
		if err != nil {
			t.Fatal(err)
		}
		if len(users) != 2 {
			t.Fatal("expected 2 users, got", len(users))
		}
		expectedUser0 := model.User{
			Group:    "group6",
			Name:     "User16Name",
			LastName: "UserLastName16",
		}
		expectedUser1 := model.User{
			Group:    "group7",
			Name:     "User17Name",
			LastName: "UserLastName17",
		}
		assertUser(t, expectedUser0, users[0])
		assertUser(t, expectedUser1, users[1])
	})
}

func TestSelectPageOrdered(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo repo.User) {
		// Select Result users
		users, err := repo.SelectPage(dbie.Page{
			Offset: 5,
			Limit:  10,
		}, `group`, dbie.In, []string{"group0", "group9"},
			dbie.Sort{Field: "group", Order: dbie.ASC},
			dbie.Sort{Field: "name", Order: dbie.DESC},
		)
		check(t, users, err).
			ExpectErr(nil).
			ExpectPageAndSize(dbie.Page{Offset: 5, Limit: 10}, 20).
			Iterate(func(prev, user model.User) {
				if user.Group != "group0" && user.Group != "group9" {
					t.Fatal("group not expected", "got", user.Group)
				}
				if prev.Group > user.Group {
					t.Fatal("invalid group order", "prev", prev, "current", user)
				}
			})
		wantFirstUser := model.User{
			Group:    "group0",
			Name:     "User40Name",
			LastName: "UserLastName40",
		}
		wantLastUser := model.User{
			Group:    "group9",
			Name:     "User69Name",
			LastName: "UserLastName69",
		}
		assertUser(t, wantFirstUser, users.Data[0])
		assertUser(t, wantLastUser, users.Data[9])
	})
}

func TestSelectNoRows(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo repo.User) {
		// Select that doesn't exist
		noUsers, err := repo.Select("last_name", dbie.Eq, "I DONT EXIST")
		if err != nil {
			t.Fatal("Select for slice should not return error but noUsers slice", err)
		}
		if len(noUsers) != 0 {
			t.Fatal("expected empty slice of users but got", noUsers)
		}
	})
}

func TestSelectError(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo repo.User) {
		data, err := repo.SelectPage(dbie.Page{}, `non existing field`, dbie.Gt, 5)
		if err == nil {
			t.Log("mongo doesn't handle schema validation well", "got", data)
		}
	})
}
func TestClose(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo repo.User) {
		err := repo.Close()
		if err != nil {
			t.Fatal("no error expected")
		}
		_ = repo.Close()
	})
}
