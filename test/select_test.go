package test

import (
	"github.com/iamgoroot/dbie"
	"testing"
)

func TestSelect(t *testing.T) {
	repo := initUserRepo(t)
	defer repo.Close()

	// Select slice
	users, err := repo.Select(`"user"."last_name"`, dbie.In, []string{"UserLastName16", "UserLastName17"})
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

func TestSelectNoRows(t *testing.T) {
	repo := initUserRepo(t)
	defer repo.Close()

	// Select that doesn't exist
	noUsers, err := repo.Select(`"user"."last_name"`, dbie.Eq, "EXPECT ERROR BECAUSE I DONT EXIST")
	if err != nil {
		t.Fatal("Select for slice should not return error but noUsers slice", err)
	}
	if noUsers != nil {
		t.Fatal("expected nil slice of users but got", noUsers)
	}
}
