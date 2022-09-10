package bun

import (
	"context"
	"fmt"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core"
	"github.com/uptrace/bun"
)

type Bun[Entity any] struct {
	context.Context
	*bun.DB
}

func New[Entity any](ctx context.Context, db *bun.DB) dbie.Repo[Entity] {
	return core.GenericBackend[Entity]{Core: Bun[Entity]{Context: ctx, DB: db}}
}

func (core Bun[Entity]) Init() error {
	var model Entity
	return core.DB.ResetModel(core.Context, &model)
}

func (core Bun[Entity]) Close() error {
	return core.DB.Close()
}

func (core Bun[Entity]) InsertCtx(ctx context.Context, items ...Entity) error {
	_, err := core.DB.NewInsert().Model(&items).Exec(ctx)
	return dbie.Wrap(err)
}

func (core Bun[Entity]) Insert(items ...Entity) error {
	return core.InsertCtx(core.Context, items...)
}

func (core Bun[Entity]) SelectPage(
	page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (dbie.Paginated[Entity], error) {
	return core.SelectPageCtx(core.Context, page, field, operator, val, orders...)
}

func (core Bun[Entity]) SelectPageCtx(
	ctx context.Context, page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (items dbie.Paginated[Entity], err error) {
	selectQuery := core.DB.NewSelect().Model(&(items.Data))
	tableName := selectQuery.GetModel().(bun.TableModel).Table().Alias
	var op string
	switch operator {
	case dbie.In:
		op = " IN (?)"
		val = bun.In(val)
	case dbie.Nin:
		op = " NOT IN (?)"
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
