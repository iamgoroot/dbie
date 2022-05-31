# dbie

## About the project

dbie (read "debie") - simple repository tool for golang with pagination

## Usage
Define model as usually in bun (for now only bun backend supported)
```golang
package model

type User struct {
	ID       int    `pg:",pk,autoincrement"`
	Name     string `pg:"name"`
	LastName string `pg:"last_name"`
	Group    string
}
```
Define wanted repository methods
```golang
package repo
import (
	"github.com/iamgoroot/dbie"
	"model"
)
//go:generate go run ../tool -type=UserRepo
type UserRepo interface {
	dbie.Repo[model.User]
	SelectByName(name string) ([]model.User, error)
	SelectByID(ID int) (model.User, error)
	SelectByGroup(page dbie.Page, group string) (items dbie.Paginated[model.User], err error)
	SelectByGroupIn(page dbie.Page, group ...string) (items dbie.Paginated[model.User], err error)
}

```

Run
```sh
go generate ./...
```
Enjoy generated implementation

## State
POC. Not stable