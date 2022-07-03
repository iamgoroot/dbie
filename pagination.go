package dbie

// Page Paginated request
type Page struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type Sort struct {
	Field string    `json:"field,omitempty"`
	Order SortOrder `json:"order,omitempty"`
}

//func (s Sort) String() string {
//	return fmt.Sprintf(`"%s" %s`, s.Field, s.Order)
//}

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
	Data   []Entity `json:"data,omitempty"`
	Offset int      `json:"offset,omitempty"`
	Limit  int      `json:"limit,omitempty"`
	Count  int      `json:"count,omitempty"`
	Order  []Sort   `json:"order,omitempty"`
}
