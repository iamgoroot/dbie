package dbie

import (
	"reflect"
	"testing"
)

func TestSelectOne(t *testing.T) {
	repo := initUserRepo(t)
	defer repo.Close()
	// Select One
	singleUser, err := repo.SelectOne(`"user"."last_name"`, Eq, "UserLastName16")
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
}

func TestSelectOneNoRows(t *testing.T) {
	repo := initUserRepo(t)
	defer repo.Close()

	// Select that doesn't exist
	notFoundUser, err := repo.SelectOne(`"user"."last_name"`, Eq, "EXPECT ERROR BECAUSE I DONT EXIST")

	// Expect dbie.ErrNoRows
	if typedErr, ok := err.(ErrNoRows); !ok {
		t.Fatal("expected error but found user", notFoundUser, "error", typedErr)
	}
	// Expect empty model
	if !reflect.DeepEqual(notFoundUser, user{}) {
		t.Fatal("expected empty user but found", notFoundUser, "error", err)
	}
}
