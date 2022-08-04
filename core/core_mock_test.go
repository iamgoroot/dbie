package core

import (
	"fmt"
	"github.com/iamgoroot/dbie"
	"reflect"
)

type CoreMock[Entity interface{}] struct {
	CloseMock      func() error
	InsertMock     func(items ...Entity) error
	InitMock       func() error
	SelectPageMock func(page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort) (items dbie.Paginated[Entity], err error)
	SelectMock     func(field string, operator dbie.Op, val any, orders ...dbie.Sort) (items []Entity, err error)
	SelectOneMock  func(field string, operator dbie.Op, val any, orders ...dbie.Sort) (item Entity, err error)
}

// Close calls CloseMock function with given fields
func (mock CoreMock[Entity]) Close() error {
	return mock.CloseMock()
}

// Insert calls InsertMock function with given fields
func (mock CoreMock[Entity]) Insert(items ...Entity) error {
	return mock.InsertMock(items...)
}

// Init calls InitMock function with given fields
func (mock CoreMock[Entity]) Init() error {
	return mock.InitMock()
}

// SelectPage calls SelectPageMock function with given fields: page, field, operator, val, wantOrders
func (mock CoreMock[Entity]) SelectPage(page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort) (items dbie.Paginated[Entity], err error) {
	return mock.SelectPageMock(page, field, operator, val, orders...)
}

// Select calls SelectMock function with given fields: field, operator, val, wantOrders
func (mock CoreMock[Entity]) Select(field string, operator dbie.Op, val any, orders ...dbie.Sort) (items []Entity, err error) {
	return mock.SelectMock(field, operator, val, orders...)
}

// SelectOne calls SelectOneMock function with given fields: field, operator, val, wantOrders
func (mock CoreMock[Entity]) SelectOne(field string, operator dbie.Op, val any, orders ...dbie.Sort) (item Entity, err error) {
	return mock.SelectOneMock(field, operator, val, orders...)
}

type returner[Mock any] func(args ...any) Mock

// SelectPageExpect sets expectations for select page
func (mock CoreMock[Entity]) SelectPageExpect(
	expectPage dbie.Page,
	expectField string,
	expectOperator dbie.Op,
	expectVal any,
	expectOrders ...dbie.Sort,
) returner[CoreMock[Entity]] {
	return func(returns ...any) CoreMock[Entity] {
		mock.SelectPageMock = func(page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort) (items dbie.Paginated[Entity], err error) {
			switch {
			case !reflect.DeepEqual(page, expectPage):
				return items, fmt.Errorf("unexpected arg: expectPage; want: %v got: %v", expectPage, page)
			case !reflect.DeepEqual(field, expectField):
				return items, fmt.Errorf("unexpected arg: field; want: %v got: %v", expectField, field)
			case !reflect.DeepEqual(operator, expectOperator):
				return items, fmt.Errorf("unexpected arg: operator; want: %v got: %v", expectOperator, operator)
			case !reflect.DeepEqual(val, expectVal):
				return items, fmt.Errorf("unexpected arg: val; want: %v got: %v", expectVal, val)
			case !reflect.DeepEqual(orders, expectOrders):
				return items, fmt.Errorf("unexpected arg: wantOrders; want: %v got: %v", expectOrders, orders)
			}
			if len(returns) > 0 && returns[0] != nil {
				items = returns[0].(dbie.Paginated[Entity])
			}
			if len(returns) > 1 && returns[1] != nil {
				err = returns[1].(error)
			}
			return
		}
		return mock
	}
}
