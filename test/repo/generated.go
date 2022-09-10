package repo

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/iamgoroot/dbie"
	coreBun "github.com/iamgoroot/dbie/core/bun"
	coreGorm "github.com/iamgoroot/dbie/core/gorm"
	corePg "github.com/iamgoroot/dbie/core/pg"
	model "github.com/iamgoroot/dbie/core/test/model"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

type UserImpl struct {
	dbie.Repo[model.User]
}

func (factory Bun) NewBunUser(ctx context.Context) User {
	return UserImpl{Repo: coreBun.New[model.User](ctx, factory.DB)}
}

func (factory Gorm) NewGormUser(ctx context.Context) User {
	return UserImpl{Repo: coreGorm.New[model.User](ctx, factory.DB)}
}

func (factory Pg) NewPgUser(ctx context.Context) User {
	return UserImpl{Repo: corePg.New[model.User](ctx, factory.DB)}
}

type Bun struct{ *bun.DB }

type Gorm struct{ *gorm.DB }

type Pg struct{ *pg.DB }

func (g UserImpl) SelectByGroup(page dbie.Page, group string) (dbie.Paginated[model.User], error) {
	return g.Repo.SelectPage(page, "group", dbie.Eq, group)
}

func (g UserImpl) SelectByGroupEq(group string) ([]model.User, error) {
	return g.Repo.Select("group", dbie.Eq, group)
}

func (g UserImpl) SelectByGroupIn(page dbie.Page, group ...string) (dbie.Paginated[model.User], error) {
	return g.Repo.SelectPage(page, "group", dbie.In, group)
}

var sortSelectByGroupNinOrderByGroupAscSetting = []dbie.Sort{
	{Field: `group`, Order: dbie.ASC},
}

func (g UserImpl) SelectByGroupNinOrderByGroupAsc(page dbie.Page, group ...string) (dbie.Paginated[model.User], error) {
	return g.Repo.SelectPage(page, "group", dbie.Nin, group, sortSelectByGroupNinOrderByGroupAscSetting...)
}

var sortSelectByGroupOrderByNameDescOrderByIDAscSetting = []dbie.Sort{
	{Field: `name`, Order: dbie.DESC},
	{Field: `id`, Order: dbie.ASC},
}

func (g UserImpl) SelectByGroupOrderByNameDescOrderByIDAsc(group string) (model.User, error) {
	return g.Repo.SelectOne("group", dbie.Eq, group, sortSelectByGroupOrderByNameDescOrderByIDAscSetting...)
}

func (g UserImpl) FindByID(id int) (model.User, error) {
	return g.Repo.SelectOne("id", dbie.Eq, id)
}

func (g UserImpl) SelectByID(id int) (model.User, error) {
	return g.Repo.SelectOne("id", dbie.Eq, id)
}

func (g UserImpl) SelectByName(name string) ([]model.User, error) {
	return g.Repo.Select("name", dbie.Eq, name)
}
