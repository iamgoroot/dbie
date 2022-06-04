package dbie

import (
	"errors"
)

type Core[Entity any] interface {
	Insert(items ...Entity) error
	SelectPage(page Page, field string, operator Op, val any, orders ...Sort) (items Paginated[Entity], err error)
}

type GenericBackend[Entity any] struct {
	Core[Entity]
}

func NewRepo[Entity any](backend Core[Entity]) Repo[Entity] {
	return GenericBackend[Entity]{
		Core: backend,
	}
}

func (p GenericBackend[Entity]) SelectOne(field string, operator Op, val any, orders ...Sort) (item Entity, err error) {
	page, err := p.SelectPage(Page{Limit: 1}, field, operator, val, orders...)
	if err == nil {
		if len(page.Data) > 0 {
			return page.Data[0], err //happy path
		}
		err = errors.New("no records found")
	}
	return
}

func (p GenericBackend[Entity]) Select(field string, operator Op, val any, orders ...Sort) (items []Entity, err error) {
	page, err := p.SelectPage(Page{}, field, operator, val, orders...)
	return page.Data, err
}
