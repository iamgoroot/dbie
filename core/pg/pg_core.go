package pg

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core"
)

type Pg[Entity any] struct {
	context.Context
	*pg.DB
}

func New[Entity any](ctx context.Context, db *pg.DB) dbie.Repo[Entity] {
	return core.GenericBackend[Entity]{Core: Pg[Entity]{Context: ctx, DB: db}}
}

func (core Pg[Entity]) Init() error {
	var model Entity
	return core.DB.Model(&model).CreateTable(&orm.CreateTableOptions{IfNotExists: true, FKConstraints: true})
}

func (core Pg[Entity]) InsertCtx(ctx context.Context, items ...Entity) error {
	_, err := core.DB.Model(&items).Context(ctx).Insert()
	return dbie.Wrap(err)
}

func (core Pg[Entity]) Insert(items ...Entity) error {
	return core.InsertCtx(core.Context, items...)
}

func (core Pg[Entity]) SelectPage(
	page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (dbie.Paginated[Entity], error) {
	return core.SelectPageCtx(core.Context, page, field, operator, val, orders...)
}

func (core Pg[Entity]) SelectPageCtx(
	ctx context.Context, page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (items dbie.Paginated[Entity], err error) {
	selectQuery := core.DB.Model(&(items.Data))
	var tName []byte
	tName, err = selectQuery.TableModel().Table().Alias.AppendValue(tName, 0)
	if err != nil {
		return
	}
	var op string
	switch operator {
	case dbie.In:
		op = " IN (?)"
		val = pg.In(val)
	case dbie.Nin:
		op = " NOT IN (?)"
		val = pg.In(val)
	default:
		op = operator.String()
	}
	tableName := string(tName)
	stmt := render(tableName, field, op)
	selectQuery.Where(stmt, val)

	for _, order := range orders {
		selectQuery.OrderExpr(render(tableName, order.Field, order.Order.String()))
	}
	selectQuery.Offset(page.Offset).Limit(page.Limit)
	var count int
	count, err = selectQuery.Context(ctx).SelectAndCount(&items.Data)
	if err != nil {
		return dbie.Paginated[Entity]{}, dbie.Wrap(err)
	}
	items.Offset, items.Limit, items.Count = page.Offset, page.Limit, count

	return items, dbie.Wrap(err)
}

func render(tableName, field, op string) string {
	return fmt.Sprintf(`%s."%s" %s`, tableName, field, op)
}
