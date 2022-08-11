package core

import (
	"github.com/iamgoroot/dbie"
)

type Core[Entity any] interface {
	Init() error
	Insert(items ...Entity) error
	SelectPage(page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort) (items dbie.Paginated[Entity], err error)
	Close() error
}

type GenericBackend[Entity any] struct {
	Core[Entity]
}

func (p GenericBackend[Entity]) Close() error {
	if p.Core != nil {
		return p.Core.Close()
	}
	return nil
}

func (p GenericBackend[Entity]) SelectOne(field string, operator dbie.Op, val any, orders ...dbie.Sort) (item Entity, err error) {
	page, err := p.SelectPage(dbie.Page{Limit: 1}, field, operator, val, orders...)
	if err == nil {
		if len(page.Data) > 0 {
			return page.Data[0], err // happy path
		}
		err = dbie.ErrNoRows
	}
	return
}

func (p GenericBackend[Entity]) Select(field string, operator dbie.Op, val any, orders ...dbie.Sort) (items []Entity, err error) {
	page, err := p.SelectPage(dbie.Page{Limit: 20}, field, operator, val, orders...)
	return page.Data, err
}
