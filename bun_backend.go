package dbie

import (
	"context"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
)

type BunBackend[Entity any] struct {
	context.Context
	*bun.DB
}

func (p BunBackend[Entity]) SelectOne(field string, operator Op, val any) (item Entity, err error) {
	items, err := p.Select(field, operator, val)
	if err != nil {
		return item, err
	}
	if len(items) == 0 {
		if err == nil {
			//TODO: err handling
			err = errors.New("no records found")
		}
		return
	}
	return items[0], err
}

func (p BunBackend[Entity]) Insert(items ...Entity) error {
	_, err := p.DB.NewInsert().Model(&items).Exec(p.Context)
	return err
}

func (p BunBackend[Entity]) Select(field string, operator Op, val any) (items []Entity, err error) {
	selectQuery := p.DB.NewSelect().Model(&items)
	switch operator {
	case In, Nin:
		selectQuery.Where(fmt.Sprint(field, operator), bun.In(val))
	default:
		selectQuery.Where(fmt.Sprint(field, operator), val)
	}
	err = selectQuery.Scan(p.Context)
	return
}

func (p BunBackend[Entity]) SelectPage(page Page, field string, operator Op, val any) (items Paginated[Entity], err error) {
	selectQuery := p.DB.NewSelect().Model(&items.Data)
	switch operator {
	case In, Nin:
		selectQuery.Where(fmt.Sprint(field, operator), bun.In(val))
	default:
		selectQuery.Where(fmt.Sprint(field, operator), val)
	}
	err = selectQuery.Offset(page.Offset).Limit(page.Limit).Scan(p.Context)
	return
}
