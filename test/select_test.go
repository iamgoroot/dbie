package test

import (
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core/test/model"
	"testing"
)

func TestSelect(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo dbie.Repo[model.User]) {
		// Select slice
		users, err := repo.Select(`last_name`, dbie.In, []string{"UserLastName16", "UserLastName17"})
		if err != nil {
			t.Fatal(err)
		}

		if len(users) != 2 {
			t.Fatal("expected 2 users, got", len(users))
		}
		// Assert
		expectedUser0 := model.User{
			Group:    "group6",
			Name:     "User16Name",
			LastName: "UserLastName16",
		}
		assertUser(t, expectedUser0, users[0])

		expectedUser1 := model.User{
			Group:    "group7",
			Name:     "User17Name",
			LastName: "UserLastName17",
		}
		assertUser(t, expectedUser1, users[1])
	})
}

func TestSelectPageOrdered(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo dbie.Repo[model.User]) {
		// Select paginated users
		users, err := repo.SelectPage(dbie.Page{
			Offset: 5,
			Limit:  10,
		}, `group`, dbie.In, []string{"group0", "group9"},
			dbie.Sort{Field: "group", Order: dbie.ASC},
			dbie.Sort{Field: "name", Order: dbie.DESC},
		)
		// Assert
		if err != nil {
			t.Fatal(err)
		}
		if err != nil {
			t.Fatal("unexpected error", err)
		}
		if users.Count != 20 {
			t.Fatal("bad result size", users.Count, users)
		}
		if users.Offset != 5 {
			t.Fatal("bad offset in result", users.Offset)
		}
		if users.Limit != 10 {
			t.Fatal("bad limit in result", users.Limit)
		}
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
	testAllCores(t, func(t *testing.T, repo dbie.Repo[model.User]) {
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
