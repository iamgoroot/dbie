package internal

import (
	"github.com/iamgoroot/dbie"
	"strings"
)

func ByDsn[Entity any](dsn string) (string, dbie.Repo[Entity]) {
	switch {
	case strings.HasPrefix(dsn, "postgresql://"), strings.HasPrefix(dsn, "postgres://"):
		return "", nil
	case strings.HasPrefix(dsn, "sqlserver://"):
		return "mssql", nil
	case strings.Contains(dsn, "@"):
		return "mysql", nil
	case strings.Contains(dsn, "file::"):
		return "sqltite", nil
	default:
		return "sqllite", nil
	}
}
