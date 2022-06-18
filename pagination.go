package dbie

import "fmt"

// Page Paginated request
type Page struct {
	Offset, Limit int
}

type Sort struct {
	Field string
	Order SortOrder
}

func (s Sort) String() string {
	return fmt.Sprintf("%s %s", s.Field, s.Order)
}

type SortOrder int

const (
	ASC = 0 + iota
	DESC
)

func (s SortOrder) String() string {
	return [...]string{"ASC", "DESC"}[s]
}

// Paginated result
type Paginated[Entity any] struct {
	Data                 []Entity
	Offset, Limit, Count int
}
