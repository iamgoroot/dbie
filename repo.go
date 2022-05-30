package dbie

type Repo[Entity any] interface {
	Insert(item ...Entity) error
	Select(field string, operator Op, val any) ([]Entity, error)
	SelectOne(field string, operator Op, val any) (Entity, error)
	SelectPage(page Page, field string, operator Op, val any) (items Paginated[Entity], err error)
}
