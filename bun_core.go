package dbie

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

type BunCore[Entity any] struct {
	GenericBackend[Entity]
	context.Context
	*bun.DB
}

func NewBun[Entity any](ctx context.Context, db *bun.DB) Repo[Entity] {
	return NewRepo[Entity](
		BunCore[Entity]{Context: ctx, DB: db},
	)
}

func (p BunCore[Entity]) InsertCtx(ctx context.Context, items ...Entity) error {
	_, err := p.DB.NewInsert().Model(&items).Exec(ctx)
	return Wrap(err)
}

func (p BunCore[Entity]) Insert(items ...Entity) error {
	return p.InsertCtx(p.Context, items...)
}

func (p BunCore[Entity]) SelectPage(
	page Page, field string, operator Op, val any, orders ...Sort,
) (Paginated[Entity], error) {
	return p.SelectPageCtx(p.Context, page, field, operator, val, orders...)
}

func (p BunCore[Entity]) SelectPageCtx(
	ctx context.Context, page Page, field string, operator Op, val any, orders ...Sort,
) (items Paginated[Entity], err error) {
	selectQuery := p.DB.NewSelect().Model(&(items.Data))
	tableName := selectQuery.GetModel().(bun.TableModel).Table().Alias
	var op string
	switch operator {
	case In:
		op = " IN (?)"
		val = bun.In(val)
	case Nin:
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
	query := selectQuery.Offset(page.Offset).Limit(page.Limit)
	var count int
	count, err = query.ScanAndCount(ctx)
	if err != nil {
		return Paginated[Entity]{}, Wrap(err)
	}
	items.Offset, items.Limit, items.Count = page.Offset, page.Limit, count

	return items, Wrap(err)
}

func render(tableName, field, op string) string {
	return fmt.Sprintf(`"%s"."%s" %s`, tableName, field, op)
}
