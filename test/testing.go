package test

import (
	"context"
	"fmt"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core/test/model"
	"github.com/iamgoroot/dbie/core/test/repo"
	"testing"
)

func testAllCores(t *testing.T, testFunc func(*testing.T, repo.User)) {
	makers := map[string]func(ctx context.Context) repo.User{
		"BunSqlite":    repo.Bun{DB: makeBunSqlite("file::memory:?")}.NewUser,
		"BunPostgres":  repo.Bun{DB: makeBunPostgres("postgres://user:pass@127.0.0.1:5432/test?sslmode=disable")}.NewUser,
		"GormSqlite":   repo.Gorm{DB: makeGormSqlite("file::memory:?")}.NewUser,
		"GormPostgres": repo.Gorm{DB: makeGormPostgres("postgres://user:pass@127.0.0.1:5433/test?sslmode=disable")}.NewUser,
		"PgPostgres":   repo.Pg{DB: makePg("postgres://user:pass@127.0.0.1:5434/test?sslmode=disable")}.NewUser,
		//TODO: "BeeGo":        repo.Bee{DB: makeGormMysql("user:pass@tcp(localhost:3307)/test")},
	}
	for key, maker := range makers {
		userRepo := maker(context.Background())
		if err := userRepo.Init(); err != nil {
			t.Fatal(err)
		}
		if err := userRepo.Insert(createUsers()...); err != nil {
			t.Fatal(err)
		}
		defer userRepo.Close()
		t.Run(key, func(t *testing.T) {
			testFunc(t, userRepo)
		})
	}
}

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

func check[T any](t *testing.T, paginated dbie.Paginated[T], err error) checker[T] {
	return checker[T]{t, paginated, err}
}

type checker[T any] struct {
	*testing.T
	Result dbie.Paginated[T]
	Err    error
}

func (e checker[T]) ExpectPageAndSize(page dbie.Page, count int) checker[T] {
	if e.Result.Offset != page.Offset {
		e.Fatal("bad offset in result", "got", e.Result.Offset, "want", page.Offset)
	}
	if e.Result.Limit != page.Limit {
		e.Fatal("bad limit in result", "got", e.Result.Limit, "want", page.Limit)
	}
	if e.Result.Count != count {
		e.Fatal("bad result size", "got", e.Result.Count, "want", count)
	}
	return e
}

func (e checker[T]) ExpectErr(err error) checker[T] {
	if e.Err != err {
		e.Fatal("bad error", "got", e.Err, "want", err)
	}
	return e
}

func (e checker[T]) Expect(f func(c checker[T])) checker[T] {
	f(e)
	return e
}
func (e checker[T]) Iterate(f func(prev T, current T)) checker[T] {
	var prev T
	for _, item := range e.Result.Data {
		f(prev, item)
		prev = item
	}
	return e
}
