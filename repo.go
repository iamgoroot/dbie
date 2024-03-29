package dbie

type Repo[Entity any] interface {
	Init() error
	Insert(item ...Entity) error
	Select(field string, operator Op, val any, orders ...Sort) ([]Entity, error)
	SelectOne(field string, operator Op, val any, orders ...Sort) (Entity, error)
	SelectPage(page Page, field string, operator Op, val any, orders ...Sort) (items Paginated[Entity], err error)
	Close() error
	// TODO: Delete(page Page, field string, operator Op, val any, orders ...Sort) (items Paginated[Entity], argErr error) ?
	// Update(items []Entity, field string, operator Op, val any) (items Paginated[Entity], argErr error) ?
}
