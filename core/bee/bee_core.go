package bee

import (
	"context"
	"github.com/beego/beego/v2/adapter/orm"
	bee "github.com/beego/beego/v2/client/orm"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core"
)

type Bee[Entity any] struct {
	core.GenericBackend[Entity]
	context.Context
	DB bee.Ormer
}

func New[Entity any](ctx context.Context, db bee.Ormer) dbie.Repo[Entity] {
	return core.NewRepo[Entity](
		Bee[Entity]{Context: ctx, DB: db},
	)
}
func (p Bee[Entity]) Init() error {
	var model Entity
	orm.RegisterModel(model)
	return nil
}

func (p Bee[Entity]) InsertCtx(ctx context.Context, items ...Entity) error {
	_, err := p.DB.InsertWithCtx(ctx, &items)
	return dbie.Wrap(err)
}

func (p Bee[Entity]) Insert(items ...Entity) error {
	return p.InsertCtx(p.Context, items...)
}

func (p Bee[Entity]) SelectPage(
	page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (dbie.Paginated[Entity], error) {
	panic("TODO")
	return p.SelectPageCtx(p.Context, page, field, operator, val, orders...)
}

func (p Bee[Entity]) SelectPageCtx(
	ctx context.Context, page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (items dbie.Paginated[Entity], err error) {
	panic("TODO")
	return items, err
}
