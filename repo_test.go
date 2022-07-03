package dbie

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func makeTestDB(t *testing.T) *bun.DB {
	connStr := "file::memory:"
	sqldb, err := sql.Open(sqliteshim.DriverName(), connStr)
	if err != nil {
		t.Fatal(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	_, _ = db.Exec("CREATE DATABASE test;")
	db.AddQueryHook(bundebug.NewQueryHook())
	_, err = db.NewCreateTable().Model(&testUser{}).Exec(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return db
}

type testUser struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	Name     string
	LastName string
	Group    string
}

func initBunRepo(t *testing.T) Repo[testUser] {
	repo := NewRepo[testUser](
		BunCore[testUser]{DB: makeTestDB(t), Context: context.Background()},
	)
	return repo
}

func makeGormTestDB(t *testing.T) *gorm.DB {
	dsn := "file::memory:"
	dialect := sqlite.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	db = db.Debug()
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(testUser{})
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func initGormRepo(t *testing.T) Repo[testUser] {
	repo := NewRepo[testUser](
		GormCore[testUser]{DB: makeGormTestDB(t), Context: context.Background()},
	)

	return repo
}

func createUsers() (users []testUser) {
	const n = 100
	users = make([]testUser, n)
	for i := 0; i < n; i++ {
		users[i] = testUser{
			Name:     fmt.Sprintf("User%dName", i),
			LastName: fmt.Sprintf("UserLastName%d", i),
			Group:    fmt.Sprintf("group%d", i%10),
		}
	}
	return users
}

func assertUser(t *testing.T, expected testUser, got testUser) {
	switch {
	case got.Group != expected.Group:
		t.Fatal("expected singleUser.Group", expected.Group, "got", got.Group)
	case got.Name != expected.Name:
		t.Fatal("expected singleUser.Name", expected.Name, "got", got.Name)
	case got.LastName != expected.LastName:
		t.Fatal("expected singleUser.LastName", expected.LastName, "got", got.LastName)
	}
}
