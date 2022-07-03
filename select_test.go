package dbie

import (
	"testing"
)

func TestSelect(t *testing.T) {
	testWithRepos(t, func(t *testing.T, repo Repo[testUser]) {
		// Select slice
		users, err := repo.Select(`last_name`, In, []string{"UserLastName16", "UserLastName17"})
		if err != nil {
			t.Fatal(err)
		}
		// Assert
		expectedUser0 := testUser{
			Group:    "group6",
			Name:     "User16Name",
			LastName: "UserLastName16",
		}
		assertUser(t, expectedUser0, users[0])

		expectedUser1 := testUser{
			Group:    "group7",
			Name:     "User17Name",
			LastName: "UserLastName17",
		}
		assertUser(t, expectedUser1, users[1])
	})
}

func TestSelectPageOrdered(t *testing.T) {
	testWithRepos(t, func(t *testing.T, repo Repo[testUser]) {
		// Select paginated users
		users, err := repo.SelectPage(Page{
			Offset: 5,
			Limit:  10,
		}, `group`, In, []string{"group0", "group9"},
			Sort{Field: "group", Order: ASC},
			Sort{Field: "name", Order: DESC},
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
		wantFirstUser := testUser{
			Group:    "group0",
			Name:     "User40Name",
			LastName: "UserLastName40",
		}
		wantLastUser := testUser{
			Group:    "group9",
			Name:     "User69Name",
			LastName: "UserLastName69",
		}

		assertUser(t, wantFirstUser, users.Data[0])
		assertUser(t, wantLastUser, users.Data[9])
	})
}

func TestSelectNoRows(t *testing.T) {
	testWithRepos(t, func(t *testing.T, repo Repo[testUser]) {
		// Select that doesn't exist
		noUsers, err := repo.Select("last_name", Eq, "I DONT EXIST")
		if err != nil {
			t.Fatal("Select for slice should not return error but noUsers slice", err)
		}
		if len(noUsers) != 0 {
			t.Fatal("expected empty slice of users but got", noUsers)
		}
	})
}
