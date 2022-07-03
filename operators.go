package dbie

type Op int

const (
	Eq = iota + 0
	Neq
	Gt
	Gte
	Lt
	Lte
	Like
	Ilike
	Nlike
	Nilike
	In
	Nin
	Is
	Not
)

func (op Op) String() string {
	return [...]string{
		" = ? ",
		" != ? ",
		" > ? ",
		" >= ? ",
		" < ? ",
		" <= ? ",
		" LIKE ?",
		" ILIKE ?",
		" NOT LIKE ?",
		" NOT ILIKE ?",
		" IN ?",
		" NOT IN ?",
		" IS NULL",
		" IS NOT NULL",
	}[op]
}
