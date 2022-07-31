package dbie

import "testing"

func TestSortOrder_String(t *testing.T) {
	tests := []struct {
		name  string
		order SortOrder
		want  string
	}{
		{
			name:  "test ASC.String()",
			order: ASC,
			want:  "ASC",
		},
		{
			name:  "test DESC.String()",
			order: DESC,
			want:  "DESC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.order.String(); got != tt.want {
				t.Errorf("String() = %s, want %s", got, tt.want)
			}
		})
	}
}
