package gormCore

import (
	"context"
	"fmt"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core"
	"gorm.io/gorm"
	"reflect"
)

type Gorm[Entity any] struct {
	context.Context
	DB *gorm.DB
}

func New[Entity any](ctx context.Context, db *gorm.DB) dbie.Repo[Entity] {
	return core.GenericBackend[Entity]{
		Gorm[Entity]{Context: ctx, DB: db},
	}
}

func (p Gorm[Entity]) Init() error {
	var model Entity
	return p.DB.WithContext(p.Context).AutoMigrate(&model)
}

func (p Gorm[Entity]) Close() error {
	if p.DB.Error != nil {
		return p.DB.Error
	}
	p.DB.AddError(fmt.Errorf("closed"))
	return nil
}

func (p Gorm[Entity]) InsertCtx(ctx context.Context, items ...Entity) error {
	err := p.DB.WithContext(ctx).Create(&items)
	return dbie.Wrap(err.Error)
}

func (p Gorm[Entity]) Insert(items ...Entity) error {
	return p.InsertCtx(p.Context, items...)
}

func (p Gorm[Entity]) SelectPage(
	page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (dbie.Paginated[Entity], error) {
	return p.SelectPageCtx(p.Context, page, field, operator, val, orders...)
}

func (p Gorm[Entity]) SelectPageCtx(
	ctx context.Context, page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (items dbie.Paginated[Entity], err error) {
	selectQuery := p.DB.Model(&(items.Data)).WithContext(ctx)
	switch operator {
	case dbie.In, dbie.Nin:
		var params []interface{}
		rv := reflect.ValueOf(val)
		if rv.Kind() == reflect.Slice {
			for i := 0; i < rv.Len(); i++ {
				params = append(params, rv.Index(i).Interface())
			}
			if err = selectQuery.Statement.Parse(&(items.Data)); err != nil {
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
		return dbie.Paginated[Entity]{}, dbie.Wrap(res.Error)
	}
	for _, order := range orders {
		selectQuery.Order(render(selectQuery.Statement.Schema.Table, order.Field, order.Order.String()))
	}
	selectQuery.Offset(page.Offset).Limit(page.Limit)
	res = selectQuery.Find(&(items.Data))
	if res.Error != nil {
		return dbie.Paginated[Entity]{}, dbie.Wrap(res.Error)
	}
	items.Count, items.Offset, items.Limit = int(count), page.Offset, page.Limit
	return items, dbie.Wrap(res.Error)
}

func render(tableName, field, op string) string {
	return fmt.Sprintf(`"%s"."%s" %s`, tableName, field, op)
}
