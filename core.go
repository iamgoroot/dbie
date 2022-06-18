package dbie

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

type BunCore[Entity any] struct {
	context.Context
	*bun.DB
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
	switch operator {
	case In, Nin:
		selectQuery.Where(fmt.Sprint(field, operator), bun.In(val))
	default:
		selectQuery.Where(fmt.Sprint(field, operator), val)
	}

	for _, order := range orders {
		selectQuery.Order(order.String())
	}
	query := selectQuery.Offset(page.Offset).Limit(page.Limit)

	items.Count, err = query.ScanAndCount(ctx)
	items.Offset, items.Limit = page.Offset, page.Limit

	return items, Wrap(err)
}
