# dbie
![Coverage](https://img.shields.io/badge/Coverage-40.7%25-yellow)
dbie (read "debie") - simple repository tool for golang with pagination

[![CI](https://github.com/iamgoroot/dbie/workflows/CI/badge.svg)](https://github.com/iamgoroot/dbie/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Update release version.](https://github.com/iamgoroot/dbie/workflows/Update%20release%20version./badge.svg)](https://github.com/iamgoroot/dbie/actions?query=workflow%3A%22Update+release+version.%22)

[![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?logo=ubuntu\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Mac OS](https://img.shields.io/badge/mac%20os-000000?logo=macos\&logoColor=F0F0F0)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Windows](https://img.shields.io/badge/Windows-0078D6?logo=windows\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)

## Usage
Define model as usually in bun (for now only bun backend supported) and wanted repository methods
```golang
package repo

import (
	"github.com/iamgoroot/dbie"
)

type User struct { //model defined as bun model. since bun is only core for now 
	ID       int    `pg:",pk,autoincrement"`
	Name     string `pg:"name"`
	LastName string `pg:"last_name"`
	Group    string
}

//go:generate go run "github.com/iamgoroot/dbietool" -type=UserRepo
type UserRepo interface {
	dbie.Repo[User]
	SelectByName(name string) ([]User, error)
	SelectByID(ID int) (User, error)
	SelectByGroup(page dbie.Page, group string) (items dbie.Paginated[User], err error)
	SelectByGroupIn(group ...string) (items dbie.Paginated[User], err error)
}

```

Run
```sh
go generate ./...
```
Enjoy generated implementation
```golang
    repo := NewRepo[User](
	BunCore[User]{DB: db, Context: context.Background()}, 
    )
    err, results = repo.SelectByGroupIn("group1", "group2")
```

# Method signature naming principles
## SelectBy
### Signatures:
* {ColumnName} - part of function name, specifically db column name but CamelCase instead of snake_case
* {?Operator} - SQL operator. 
  * `dbie.Eq` if omitted. 
  * Possible values:
  `"Eq" (default), "Neq", "Gt", "Gte", "Lt", "Lte", "Like", "Ilike", "Nlike", "Nilike", "In", "Nin", "Is", "Not"`
* {columnName} - columnName in camelCase.
* {columnType} - type of parameter as golang type
* Supported return types:
  * MODEL - returns one item 
  * []MODEL - returns slice of resulting items
  * dbie.Paginated[MODEL] - returns paginated wrapper with resulting items
* Each method returns error as second parameter

```golang
func SelectBy{ColumnName}({columnName} {columnType}) (MODEL, error) // returns one row or error 
func SelectBy{ColumnName}{?Operator}( {columnName} {columnType} ) (MODEL, error) // returns one row or error 
func SelectBy{ColumnName}{?Operator}( {columnName} {columnType} ) ([]MODEL, error) // returns slice or error
func SelectBy{ColumnName}{?Operator}( {columnName} {columnType} ) (dbie.Paginated[MODEL], error) // returns slice wrapper with pagination or error
```




### OrderBy patterns:

* {OrderColumnName} - ColumnName to order by in CamelCase.
* {?SortOrder} - Asc or Desc
* columnName and columnType as in previous example
* composite sorting is supported
```golang
func SelectByColumnNameOrderBy{OrderColumnName}{?SortOrder}(columnName columnType) ([]MODEL, error)
func SelectByColumnNameOrderBy{OrderColumnName}{?SortOrder}{ColumnName2}{?Order2}(columnName columnType) ([]MODEL, error)

```
