package dbie

// Page Paginated request
type Page struct {
	Offset, Limit int
	Sort          string
}

//Paginated result
type Paginated[Entity any] struct {
	Data                 []Entity
	Offset, Limit, Total int
}
