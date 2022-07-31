package test

import (
	"fmt"
	"github.com/iamgoroot/dbie/core/test/model"
	"testing"
)

func createUsers() (users []model.User) {
	const n = 100
	users = make([]model.User, n)
	for i := 0; i < n; i++ {
		users[i] = model.User{
			Name:     fmt.Sprintf("User%dName", i),
			LastName: fmt.Sprintf("UserLastName%d", i),
			Group:    fmt.Sprintf("group%d", i%10),
		}
	}
	return users
}

func assertUser(t *testing.T, expected model.User, got model.User) {
	switch {
	case got.Group != expected.Group:
		t.Fatal("expected singleUser.Group", expected.Group, "got", got.Group)
	case got.Name != expected.Name:
		t.Fatal("expected singleUser.Name", expected.Name, "got", got.Name)
	case got.LastName != expected.LastName:
		t.Fatal("expected singleUser.LastName", expected.LastName, "got", got.LastName)
	}
}
