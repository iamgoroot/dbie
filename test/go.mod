module github.com/iamgoroot/dbie/core/test

go 1.18

replace (
	github.com/iamgoroot/dbie => ./..
	github.com/iamgoroot/dbie/core/bee => ./../core/bee
	github.com/iamgoroot/dbie/core/bun => ./../core/bun
	github.com/iamgoroot/dbie/core/gorm => ./../core/gorm
)

require (
	github.com/iamgoroot/dbie v0.0.0-20220715232405-9fcd7d479f61
	github.com/iamgoroot/dbie/core/bun v0.0.0-00010101000000-000000000000
	github.com/iamgoroot/dbie/core/gorm v0.0.0-00010101000000-000000000000
	github.com/uptrace/bun v1.1.6
	github.com/uptrace/bun/dialect/mysqldialect v1.1.6
	github.com/uptrace/bun/dialect/pgdialect v1.1.6
	github.com/uptrace/bun/dialect/sqlitedialect v1.1.6
	github.com/uptrace/bun/driver/pgdriver v1.1.6
	github.com/uptrace/bun/driver/sqliteshim v1.1.6
	github.com/uptrace/bun/extra/bundebug v1.1.6
	gorm.io/driver/mysql v1.3.5
	gorm.io/driver/postgres v1.3.8
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/lib/pq v1.10.5 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v1.14.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220708085239-5a0f0661e09d // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.11 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	lukechampine.com/uint128 v1.2.0 // indirect
	mellium.im/sasl v0.2.1 // indirect
	modernc.org/cc/v3 v3.36.0 // indirect
	modernc.org/ccgo/v3 v3.16.7 // indirect
	modernc.org/libc v1.16.14 // indirect
	modernc.org/mathutil v1.4.1 // indirect
	modernc.org/memory v1.1.1 // indirect
	modernc.org/opt v0.1.3 // indirect
	modernc.org/sqlite v1.17.3 // indirect
	modernc.org/strutil v1.1.2 // indirect
	modernc.org/token v1.0.0 // indirect
)
