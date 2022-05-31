package dbie

import (
	"errors"
)

type Backend[Entity any] interface {
	Insert(items ...Entity) error
	SelectPage(page Page, field string, operator Op, val any, orders ...Sort) (items Paginated[Entity], err error)
}
type BaseBackend[Entity any] struct {
	Backend[Entity]
}

func NewRepo[Entity any](backend Backend[Entity]) Repo[Entity] {
	return BaseBackend[Entity]{
		Backend: backend,
	}
}

func (p BaseBackend[Entity]) SelectOne(field string, operator Op, val any) (item Entity, err error) {
	page, err := p.SelectPage(Page{Limit: 1}, field, operator, val)
	if err == nil {
		if len(page.Data) > 0 {
			return page.Data[0], err //happy path
		}
		err = errors.New("no records found")
	}
	return
}

func (p BaseBackend[Entity]) Select(field string, operator Op, val any) (items []Entity, err error) {
	page, err := p.SelectPage(Page{}, field, operator, val)
	return page.Data, err
}
