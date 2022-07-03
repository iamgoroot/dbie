package dbie

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

type GormCore[Entity any] struct {
	GenericBackend[Entity]
	context.Context
	DB *gorm.DB
}

func (p GormCore[Entity]) Close() error {
	db, err := p.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (p GormCore[Entity]) InsertCtx(ctx context.Context, items ...Entity) error {
	err := p.DB.WithContext(ctx).Create(&items)
	return Wrap(err.Error)
}

func (p GormCore[Entity]) Insert(items ...Entity) error {
	return p.InsertCtx(p.Context, items...)
}

func (p GormCore[Entity]) SelectPage(
	page Page, field string, operator Op, val any, orders ...Sort,
) (Paginated[Entity], error) {
	return p.SelectPageCtx(p.Context, page, field, operator, val, orders...)
}

func (p GormCore[Entity]) SelectPageCtx(
	ctx context.Context, page Page, field string, operator Op, val any, orders ...Sort,
) (items Paginated[Entity], err error) {
	selectQuery := p.DB.Model(&(items.Data)).WithContext(ctx)
	switch operator {
	case In, Nin:
		var params []interface{}
		rv := reflect.ValueOf(val)
		if rv.Kind() == reflect.Slice {
			for i := 0; i < rv.Len(); i++ {
				params = append(params, rv.Index(i).Interface())
			}
			err = selectQuery.Statement.Parse(&(items.Data))
			if err != nil {
				return
			}
			whereStmt := fmt.Sprintf(`"%s"."%s" %s`, selectQuery.Statement.Schema.Table, field, operator)
			selectQuery = selectQuery.Where(whereStmt, params)
			break
		}
		fallthrough
	default:
		selectQuery = selectQuery.Where(fmt.Sprint(field, operator), val)
	}
	var count int64
	res := selectQuery.Count(&count)
	if res.Error != nil {
		return Paginated[Entity]{}, Wrap(res.Error)
	}
	for _, order := range orders {
		selectQuery.Order(render(selectQuery.Statement.Schema.Table, order.Field, order.Order.String()))
	}
	selectQuery.Offset(page.Offset).Limit(page.Limit)
	res = selectQuery.Find(&(items.Data))
	if res.Error != nil {
		return Paginated[Entity]{}, Wrap(res.Error)
	}
	items.Count, items.Offset, items.Limit = int(count), page.Offset, page.Limit
	return items, Wrap(res.Error)
}
