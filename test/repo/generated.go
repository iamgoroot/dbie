package repo

import (
	bunCore "github.com/iamgoroot/dbie/core/bun"
	gormCore "github.com/iamgoroot/dbie/core/gorm"
	"github.com/iamgoroot/dbie/core/test/model"
	"github.com/uptrace/bun"
	"gorm.io/gorm"

	"context"
	"github.com/iamgoroot/dbie"
)

type UserImpl struct {
	dbie.Repo[model.User]
}

type Bun struct {
	*bun.DB
}

func (db Bun) NewUser(ctx context.Context) User {
	return UserImpl{Repo: bunCore.New[model.User](ctx, db.DB)}
}

type Gorm struct {
	*gorm.DB
}

func (db Gorm) NewUser(ctx context.Context) User {
	return UserImpl{Repo: gormCore.New[model.User](ctx, db.DB)}
}

func (g UserImpl) SelectByName(name string) ([]model.User, error) {
	return g.Repo.Select("name", dbie.Eq, name)
}

func (g UserImpl) SelectByGroupEq(group string) ([]model.User, error) {
	return g.Repo.Select("group", dbie.Eq, group)
}

func (g UserImpl) SelectByGroup(page dbie.Page, group string) (dbie.Paginated[model.User], error) {
	return g.Repo.SelectPage(page, "group", dbie.Eq, group)
}

func (g UserImpl) SelectByGroupIn(page dbie.Page, group ...string) (dbie.Paginated[model.User], error) {
	return g.Repo.SelectPage(page, "group", dbie.In, group)
}

var sortSelectByGroupOrderByNameDescOrderByIDAscSetting = []dbie.Sort{
	{Field: `name`, Order: dbie.DESC},
	{Field: `id`, Order: dbie.ASC},
}

func (g UserImpl) SelectByGroupOrderByNameDescOrderByIDAsc(group string) (model.User, error) {
	return g.Repo.SelectOne("group", dbie.Eq, group, sortSelectByGroupOrderByNameDescOrderByIDAscSetting...)
}

func (g UserImpl) SelectByID(id int) (model.User, error) {
	return g.Repo.SelectOne("id", dbie.Eq, id)
}

func (g UserImpl) FindByID(id int) (model.User, error) {
	return g.Repo.SelectOne("id", dbie.Eq, id)
}
