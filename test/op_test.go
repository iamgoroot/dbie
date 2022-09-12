package test

import (
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core/test/repo"
	"testing"
)

func TestOperators(t *testing.T) {
	type testData struct {
		Name      string
		Count     int
		Page      dbie.Page
		QueryArgs any
	}
	makeTestData := func(name string, count int, page dbie.Page, queryArgs any) testData {
		return testData{
			Name:      name,
			Count:     count,
			Page:      page,
			QueryArgs: queryArgs,
		}
	}
	ops := map[dbie.Op]testData{
		dbie.Eq:     makeTestData("test equal", 1, dbie.Page{}, "User0Name"),
		dbie.Neq:    makeTestData("test not equal", 99, dbie.Page{}, "User0Name"),
		dbie.Gt:     makeTestData("test greater", 99, dbie.Page{}, "User0Name"),
		dbie.Gte:    makeTestData("test greater or equal", 100, dbie.Page{}, "User0Name"),
		dbie.Lt:     makeTestData("test less", 0, dbie.Page{}, "User0Name"),
		dbie.Lte:    makeTestData("test less or equal", 1, dbie.Page{}, "User0Name"),
		dbie.Like:   makeTestData("test like", 1, dbie.Page{}, "User0Name"),
		dbie.Ilike:  makeTestData("test ilike", 1, dbie.Page{}, "user0Name"),
		dbie.Nlike:  makeTestData("test not like", 99, dbie.Page{}, "User0Name"),
		dbie.Nilike: makeTestData("test not ilike", 99, dbie.Page{Limit: 90}, "User0name"),
		dbie.In:     makeTestData("test in", 1, dbie.Page{}, []string{"User0Name"}),
		dbie.Nin:    makeTestData("test not in", 99, dbie.Page{}, []string{"User0Name"}),
		dbie.Is:     makeTestData("test is null", 0, dbie.Page{}, ""),
		dbie.Not:    makeTestData("test is not null. ignore args", 100, dbie.Page{}, "User0Name"),
	}
	testAllCores(t, func(t *testing.T, repo repo.User) {
		for op, expect := range ops {
			t.Run(expect.Name, func(t *testing.T) {
				res, err := repo.SelectPage(expect.Page, "name", op, expect.QueryArgs)
				check(t, res, err).ExpectPageAndSize(expect.Page, expect.Count)

			})
		}
	})
}
