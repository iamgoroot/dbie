package bun

import (
	"context"
	"fmt"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core"
	"github.com/uptrace/bun"
)

type Bun[Entity any] struct {
	core.GenericBackend[Entity]
	context.Context
	*bun.DB
}

func New[Entity any](ctx context.Context, db *bun.DB) dbie.Repo[Entity] {
	return core.NewRepo[Entity](
		Bun[Entity]{Context: ctx, DB: db},
	)
}

func (p Bun[Entity]) Init() error {
	var model Entity
	return p.DB.ResetModel(p.Context, &model)
}

func (p Bun[Entity]) InsertCtx(ctx context.Context, items ...Entity) error {
	_, err := p.DB.NewInsert().Model(&items).Exec(ctx)
	return dbie.Wrap(err)
}

func (p Bun[Entity]) Insert(items ...Entity) error {
	return p.InsertCtx(p.Context, items...)
}

func (p Bun[Entity]) SelectPage(
	page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (dbie.Paginated[Entity], error) {
	return p.SelectPageCtx(p.Context, page, field, operator, val, orders...)
}

func (p Bun[Entity]) SelectPageCtx(
	ctx context.Context, page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (items dbie.Paginated[Entity], err error) {
	selectQuery := p.DB.NewSelect().Model(&(items.Data))
	tableName := selectQuery.GetModel().(bun.TableModel).Table().Alias
	var op string
	switch operator {
	case dbie.In:
		op = " IN (?)"
		val = bun.In(val)
	case dbie.Nin:
		op = " NIN (?)"
		val = bun.In(val)
	default:
		op = operator.String()
	}
	stmt := render(tableName, field, op)
	selectQuery.Where(stmt, val)

	for _, order := range orders {
		selectQuery.OrderExpr(render(tableName, order.Field, order.Order.String()))
	}
	selectQuery.Offset(page.Offset).Limit(page.Limit)
	var count int
	count, err = selectQuery.ScanAndCount(ctx)
	if err != nil {
		return dbie.Paginated[Entity]{}, dbie.Wrap(err)
	}
	items.Offset, items.Limit, items.Count = page.Offset, page.Limit, count

	return items, dbie.Wrap(err)
}

func render(tableName, field, op string) string {
	return fmt.Sprintf(`"%s"."%s" %s`, tableName, field, op)
}
