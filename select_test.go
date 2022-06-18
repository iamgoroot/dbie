package dbie

import (
	"testing"
)

func TestSelect(t *testing.T) {
	repo := initUserRepo(t)
	defer repo.Close()

	// Select slice
	users, err := repo.Select(`"user"."last_name"`, In, []string{"UserLastName16", "UserLastName17"})
	if err != nil {
		t.Fatal(err)
	}
	// Assert
	expectedUser0 := user{
		Group:    "group6",
		Name:     "User16Name",
		LastName: "UserLastName16",
	}
	assertUser(t, expectedUser0, users[0])

	expectedUser1 := user{
		Group:    "group7",
		Name:     "User17Name",
		LastName: "UserLastName17",
	}
	assertUser(t, expectedUser1, users[1])
}

func TestSelectPageOrdered(t *testing.T) {
	repo := initUserRepo(t)
	defer repo.Close()

	// Select paginated users
	users, err := repo.SelectPage(Page{
		Offset: 5,
		Limit:  10,
	}, `"user"."group"`, In, []string{"group0", "group9"},
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
		t.Fatal("bad result size", users.Count)
	}
	if users.Offset != 5 {
		t.Fatal("bad offset in result", users.Offset)
	}
	if users.Limit != 10 {
		t.Fatal("bad limit in result", users.Limit)
	}
	wantFirstUser := user{
		Group:    "group0",
		Name:     "User40Name",
		LastName: "UserLastName40",
	}
	wantLastUser := user{
		Group:    "group9",
		Name:     "User69Name",
		LastName: "UserLastName69",
	}

	assertUser(t, wantFirstUser, users.Data[0])
	assertUser(t, wantLastUser, users.Data[9])
}

func TestSelectNoRows(t *testing.T) {
	repo := initUserRepo(t)
	defer repo.Close()

	// Select that doesn't exist
	noUsers, err := repo.Select(`"user"."last_name"`, Eq, "I DONT EXIST")
	if err != nil {
		t.Fatal("Select for slice should not return error but noUsers slice", err)
	}
	if len(noUsers) != 0 {
		t.Fatal("expected empty slice of users but got", noUsers)
	}
}
