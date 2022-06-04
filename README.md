# dbie

## About the project

dbie (read "debie") - simple repository tool for golang with pagination

## Usage
Define model as usually in bun (for now only bun backend supported) and wanted repository methods
```golang
package repo

import (
	"github.com/iamgoroot/dbie"
)

type User struct {
	ID       int    `pg:",pk,autoincrement"`
	Name     string `pg:"name"`
	LastName string `pg:"last_name"`
	Group    string
}

//go:generate go run "github.com/iamgoroot/dbietool" -type=UserRepo
type UserRepo interface {
	dbie.Repo[User]
	SelectByName(name string) ([]User, error)
	SelectByID(ID int) (User, error)
	SelectByGroup(page dbie.Page, group string) (items dbie.Paginated[User], err error)
	SelectByGroupIn(page dbie.Page, group ...string) (items dbie.Paginated[User], err error)
}

```

Run
```sh
go generate ./...
```
Enjoy generated implementation

## State
POC. Not stable