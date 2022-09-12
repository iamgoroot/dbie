package mongo

import (
	"context"
	"github.com/iamgoroot/dbie"
	"github.com/iamgoroot/dbie/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type Mongo[Entity any] struct {
	context.Context
	DB             *mongo.Database
	CollectionName string
}

func New[Entity any](ctx context.Context, db *mongo.Database) dbie.Repo[Entity] {
	typeInfo := reflect.TypeOf((*Entity)(nil))
	collectionName := typeInfo.Elem().Name()
	return core.GenericBackend[Entity]{Core: Mongo[Entity]{Context: ctx, DB: db, CollectionName: collectionName}}
}

func (core Mongo[Entity]) Init() error {
	return nil
}

func (core Mongo[Entity]) Close() error {
	return core.DB.Client().Disconnect(core.Context)
}

func (core Mongo[Entity]) InsertCtx(ctx context.Context, items ...Entity) error {
	core.collection().InsertMany(ctx, core.conv(items...))
	return nil
}

func (core Mongo[Entity]) Insert(items ...Entity) error {
	return core.InsertCtx(core.Context, items...)
}

func (core Mongo[Entity]) SelectPage(
	page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (dbie.Paginated[Entity], error) {
	return core.SelectPageCtx(core.Context, page, field, operator, val, orders...)
}

func (core Mongo[Entity]) SelectPageCtx(
	ctx context.Context, page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort,
) (items dbie.Paginated[Entity], err error) {
	sort := make(bson.D, len(orders))
	for i, order := range orders {
		sort[i] = bson.E{Key: order.Field, Value: convOrder(order.Order)}
	}
	opts := options.Find().SetSort(sort).SetLimit(int64(page.Limit)).SetSkip(int64(page.Offset))
	countOpts := options.Count()
	var cursor *mongo.Cursor
	collection := core.collection()
	var count int64
	filter := op(operator, field, val)
	count, err = collection.CountDocuments(ctx, filter, countOpts)
	if err != nil {
		return dbie.Paginated[Entity]{}, err
	}
	cursor, err = collection.Find(ctx, filter, opts)
	if err != nil {
		return dbie.Paginated[Entity]{}, err
	}
	if err = cursor.All(ctx, &items.Data); err != nil {
		return dbie.Paginated[Entity]{}, err
	}
	items.Count, items.Offset, items.Limit = int(count), page.Offset, page.Limit
	return
}

func (core Mongo[Entity]) collection() *mongo.Collection {
	return core.DB.Collection(core.CollectionName)
}

func (core Mongo[Entity]) conv(entities ...Entity) []interface{} {
	res := make([]interface{}, len(entities))
	for i, e := range entities {
		res[i] = e
	}
	return res
}

func op(dbOp dbie.Op, field string, val any) any {
	switch dbOp {
	case dbie.Eq:
		return bson.D{{field, val}}
	case dbie.Neq:
		return bson.D{{field, bson.M{"$ne": val}}}
	case dbie.Gt:
		return bson.D{{field, bson.M{"$gt": val, "$exists": true}}}
	case dbie.Gte:
		return bson.D{{field, bson.D{{"$gte", val}}}}
	case dbie.Lt:
		return bson.D{{field, bson.D{{"$lt", val}}}}
	case dbie.Lte:
		return bson.D{{field, bson.D{{"$lte", val}}}}
	case dbie.Like:
		return bson.D{{field, bson.M{"$regex": val}}}
	case dbie.Ilike:
		return bson.D{{field, bson.M{"$regex": val, "$options": "i"}}}
	case dbie.Nlike:
		return bson.D{{field, bson.M{"$not": bson.M{"$regex": val}}}}
	case dbie.Nilike:
		return bson.D{{field, bson.M{"$not": bson.M{"$regex": val, "$options": "i"}}}}
	case dbie.In:
		return bson.M{field: bson.M{"$in": val, "$exists": true}}
	case dbie.Nin:
		return bson.M{field: bson.M{"$nin": val, "$exists": true}}
	case dbie.Is:
		return bson.D{{field, bson.M{"$exists": false}}}
	case dbie.Not:
		return bson.D{{field, bson.M{"$exists": true}}}
	default:
		return bson.D{}
	}

}

func convOrder(order dbie.SortOrder) int {
	if order == dbie.ASC {
		return 1
	}
	return -1
}
