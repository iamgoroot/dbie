package test

import (
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core/test/model"
	"github.com/iamgoroot/dbie/core/test/repo"
	"testing"
)

func TestGeneratedMethods(t *testing.T) {
	testAllCores(t, func(t *testing.T, repo repo.User) {
		users, err := repo.SelectByGroupIn(dbie.Page{Offset: 5, Limit: 10}, "group0", "group9")
		check(t, users, err).
			ExpectErr(nil).
			ExpectPageAndSize(dbie.Page{Offset: 5, Limit: 10}, 20).
			Iterate(func(_, user model.User) {
				if user.Group != "group0" && user.Group != "group9" {
					t.Fatal("group not expected", "got", user.Group)
				}
			})
		users, err = repo.SelectByGroupNinOrderByGroupAsc(dbie.Page{Offset: 1, Limit: 70}, "group0", "group1")
		check(t, users, err).
			ExpectErr(nil).
			ExpectPageAndSize(dbie.Page{Offset: 1, Limit: 70}, 80).
			Iterate(func(prev, user model.User) {
				if user.Group == "group0" || user.Group == "group1" {
					t.Fatal("group not expected", "got", user.Group)
				}
				if prev.Group > user.Group {
					t.Fatal("invalid group order", "prev", prev, "current", user)
				}
			})
	})
}
