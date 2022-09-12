module github.com/iamgoroot/dbie/core/test

go 1.18

replace (
	github.com/iamgoroot/dbie => ./..
	github.com/iamgoroot/dbie/core/bun => ./../core/bun
	github.com/iamgoroot/dbie/core/gorm => ./../core/gorm
	github.com/iamgoroot/dbie/core/mongo => ./../core/mongo
	github.com/iamgoroot/dbie/core/pg => ./../core/pg
)

require (
	github.com/go-pg/pg/extra/pgdebug/v10 v10.10.6
	github.com/go-pg/pg/v10 v10.10.6
	github.com/iamgoroot/dbie v0.0.0-00010101000000-000000000000
	github.com/iamgoroot/dbie/core/bun v0.0.0-00010101000000-000000000000
	github.com/iamgoroot/dbie/core/gorm v0.0.0-00010101000000-000000000000
	github.com/iamgoroot/dbie/core/mongo v0.0.0-00010101000000-000000000000
	github.com/iamgoroot/dbie/core/pg v0.0.0-00010101000000-000000000000
	github.com/uptrace/bun v1.1.8
	github.com/uptrace/bun/dialect/pgdialect v1.1.8
	github.com/uptrace/bun/dialect/sqlitedialect v1.1.8
	github.com/uptrace/bun/driver/pgdriver v1.1.8
	github.com/uptrace/bun/driver/sqliteshim v1.1.8
	github.com/uptrace/bun/extra/bundebug v1.1.8
	go.mongodb.org/mongo-driver v1.10.2
	gorm.io/driver/postgres v1.3.9
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-pg/zerochecker v0.2.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/iamgoroot/dbietool v0.0.0-20220912013132-99d8e85467f5 // indirect
	github.com/iamgoroot/merge v0.0.0-20220619000559-5331410dde5e // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
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
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/bufpool v0.1.11 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20220826181053-bd7e27e6170d // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.0.0-20220909162455-aba9fc2a8ff2 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	lukechampine.com/uint128 v1.2.0 // indirect
	mellium.im/sasl v0.3.0 // indirect
	modernc.org/cc/v3 v3.36.3 // indirect
	modernc.org/ccgo/v3 v3.16.9 // indirect
	modernc.org/libc v1.17.1 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.2.1 // indirect
	modernc.org/opt v0.1.3 // indirect
	modernc.org/sqlite v1.18.1 // indirect
	modernc.org/strutil v1.1.3 // indirect
	modernc.org/token v1.0.1 // indirect
)
