package dbie

import "testing"

func TestOp_String(t *testing.T) {
	tests := []struct {
		name string
		op   Op
		want string
	}{
		{
			name: "test Eq.String()",
			op:   Eq,
			want: " = ? ",
		},
		{
			name: "test Neq.String()",
			op:   Neq,
			want: " != ? ",
		},
		{
			name: "test Gt.String()",
			op:   Gt,
			want: " > ? ",
		},
		{
			name: "test Gte.String()",
			op:   Gte,
			want: " >= ? ",
		},
		{
			name: "test Lt.String()",
			op:   Lt,
			want: " < ? ",
		},
		{
			name: "test Lte.String()",
			op:   Lte,
			want: " <= ? ",
		},
		{
			name: "test Like.String()",
			op:   Like,
			want: " LIKE ?",
		},
		{
			name: "test Ilike.String()",
			op:   Ilike,
			want: " ILIKE ?",
		},
		{
			name: "test Nlike.String()",
			op:   Nlike,
			want: " NOT LIKE ?",
		},
		{
			name: "test Nilike.String()",
			op:   Nilike,
			want: " NOT ILIKE ?",
		},
		{
			name: "test In.String()",
			op:   In,
			want: " IN ?",
		},
		{
			name: "test Nin.String()",
			op:   Nin,
			want: " NOT IN ?",
		},
		{
			name: "test Is.String()",
			op:   Is,
			want: " IS NULL",
		},
		{
			name: "test Not.String()",
			op:   Not,
			want: " IS NOT NULL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.op.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
