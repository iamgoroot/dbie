package dbie

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

type BunBackend[Entity any] struct {
	Backend[Entity]
	context.Context
	*bun.DB
}

func (p BunBackend[Entity]) Insert(items ...Entity) error {
	_, err := p.DB.NewInsert().Model(&items).Exec(p.Context)
	return err
}

func (p BunBackend[Entity]) SelectPage(page Page, field string, operator Op, val any, orders ...Sort) (items Paginated[Entity], err error) {
	selectQuery := p.DB.NewSelect().Model(&items.Data)
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
	items.Count, err = query.ScanAndCount(p.Context)
	return
}
