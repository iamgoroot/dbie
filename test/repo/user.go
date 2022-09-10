package repo

import (
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core/test/model"
)

//go:generate dbietool -core=Bun,Gorm,Pg -constr=factory

type User interface {
	dbie.Repo[model.User]
	Init() error
	SelectByName(string) ([]model.User, error)
	SelectByID(int) (model.User, error)
	FindByID(int) (model.User, error)
	SelectByGroupEq(string) ([]model.User, error)
	SelectByGroup(dbie.Page, string) (items dbie.Paginated[model.User], err error)
	SelectByGroupIn(dbie.Page, ...string) (items dbie.Paginated[model.User], err error)
	SelectByGroupNinOrderByGroupAsc(dbie.Page, ...string) (items dbie.Paginated[model.User], err error)
	SelectByGroupOrderByNameDescOrderByIDAsc(string) (model.User, error)
}
