# dbie

[![codecov](https://codecov.io/gh/iamgoroot/dbie/branch/main/graph/badge.svg?token=HDGXEOT8BA)](https://codecov.io/gh/iamgoroot/dbie)

dbie - (DB Interface Extension) generates database layer implementation by simply defining its interface.

0. [Why?](#Why?)
1. [Usage](#Usage)
   1. [Define database model](#Define database model)
   2. [Define repo interface](#Define repo interface)
   3. [Generate repository implementation](#Generate repository implementation)
   4. [Use generated repository](#Use generated repository)
2. [Naming convention](#Naming convention)
   1. [SelectBy*|FindBy*](#SelectBy*|FindBy*)
3. [Custom methods](#Custom methods)

## Why?
Easy:
   * You provide contract in form on an interface 
   * Dbie provides an implementation for methods that match naming convention

## Usage

 ### Define database model
As usually in bun, gorm or beego:
```golang
type User struct {
	ID       int
	Name     string
	Group    string
}
```
 ### Define repo interface
Define methods you want implemented by using [naming convention](#Naming convention) and use
wrappers for pagination (`dbie.Page` and `dbie.Paginated`)

```golang 
// go:generate go run "github.com/iamgoroot/dbietool" -core=Bun,Gorm
type UserRepo interface {
	dbie.Repo[User] // add basic repo methods if needed
	SelectByName(name string) ([]User, error)
	SelectByID(ID int) (User, error)
	FindBy(ID int) (User, error)
	SelectByGroup(page dbie.Page, group string) (items dbie.Paginated[User], err error)
	SelectByGroupIn(group ...string) (items dbie.Paginated[User], err error)
}

```

 ### Generate repository implementation

```sh
go generate ./...
```

 ### Use generated repository
Enjoy generated implementation with repository factory (by default)
```golang
factory := repo.Bun{DB: db}
r := factory.NewUser(context.Background())
r.SelectByName("John")
```
Or run dbietool with flag `-constr=func` to use function constructors instead of factory

```golang
r := repo.NewUser(context.Background())
r.SelectByName("John")
```


# Naming convention

## SelectBy*|FindBy*

Can be used to select items by some criteria.

### Criteria
For now only one criteria is supported per method. 
For more complex queries create [custom method](#Custom methods) implementation

```golang
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
func FindBy{ColumnName}({columnName} {columnType}) (MODEL, error) // same as above
func SelectBy{ColumnName}{?Operator}( {columnName} {columnType} ) (MODEL, error) // returns one row or error 
func SelectBy{ColumnName}{?Operator}( {columnName} {columnType} ) ([]MODEL, error) // returns slice or error
func SelectBy{ColumnName}{?Operator}( {columnName} {columnType} ) (dbie.Paginated[MODEL], error) // returns slice wrapper with pagination or error
```




### Ordering:

* {OrderColumnName} - ColumnName to order by in CamelCase.
* {?SortOrder} - Asc or Desc
* columnName and columnType as in previous example
* composite sorting is supported
```golang
func SelectByColumnNameOrderBy{OrderColumnName}{?SortOrder}(columnName columnType) ([]MODEL, error)
func SelectByColumnNameOrderBy{OrderColumnName}{?SortOrder}{ColumnName2}{?Order2}(columnName columnType) ([]MODEL, error)

```


# Custom methods

1. Create separate file in same package as repo implementation
2. Create method with desired signature that does start with SelectBy* or FindBy*
