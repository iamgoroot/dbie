package dbie

import "fmt"

// Page Paginated request
type Page struct {
	Offset, Limit int
}

type Sort struct {
	Field string
	Desc  bool
}

func (s Sort) String() string {
	if s.Desc {
		return fmt.Sprintf("%s DESC", s.Field)
	}
	return fmt.Sprintf("%s ASC", s.Field)
}

//Paginated result
type Paginated[Entity any] struct {
	Data                 []Entity
	Offset, Limit, Total int
}
