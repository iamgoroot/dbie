# dbie

[![codecov](https://codecov.io/gh/iamgoroot/dbie/branch/main/graph/badge.svg?token=HDGXEOT8BA)](https://codecov.io/gh/iamgoroot/dbie)

dbie - (DB Interface Extension) generates database layer implementation by simply defining its interface.

1. [Why?](#why?)
2. [Getting started](#getting-started)
   1. [Install](#install-generator-tool)
   2. [Define contracts](#define-contracts)
   3. [Usage](#Usage)
3. [Naming convention](#naming-convention)
   1. [SelectBy*|FindBy*](#SelectBy*|FindBy*)
   2. [Sort order](#sort-order)
4. [Custom methods](#custom-methods)

## Why?
   * You provide contract in form on an interface 
   * Dbie provides an implementation for methods matching signature convention

## Getting started

### Install generator tool
```sh
   go get -u github.com/iamgoroot/dbietool
   go install github.com/iamgoroot/dbietool
```
### Define contracts
#### Define model

As usually in bun, gorm or pg:
```golang
type User struct {
	ID       int
	Name     string
	Group    string
}
```
#### Define repository interface
Define methods you want implemented by using [naming convention](#Naming convention) and use
wrappers for pagination (`dbie.Page` and `dbie.Paginated`)

```golang 
//go:generate dbietool -core=Bun,Gorm,Pg -constr=factory

type User interface {
	dbie.Repo[model.User]
	Init() error
	SelectByName(string) ([]model.User, error)
	SelectByID(int) (model.User, error)
	FindByID(int) (model.User, error)
	SelectByGroupEq(string) ([]model.User, error)
	SelectByGroup(dbie.Page, string) (items dbie.Paginated[model.User], err error)
	SelectByGroupIn(dbie.Page, ...string) (items dbie.Paginated[model.User], err error)
	SelectByGroupNinOrderByGroupAsc(dbie.Page, ...string) (items dbie.Paginated[model.User], err error)
	SelectByGroupOrderByNameDescOrderByIDAsc(string) (model.User, error)
}
```

### Generate
That's it. generate code
   ```sh
   go generate ./...
   ```

### Usage

```
func main() {
	// instantiate (run dbietool with `-constr=func` parameter)
	userRepo := repo.NewUser(context.Background())
	
	// insert user and handle error
	err := userRepo.Insert(model.User{Name: "userName1"})
	if err != nil {
		log.Fatalln(err)
	}
	
	// select user using generated method and handle error
	user, err := userRepo.SelectByName("userName1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(user, err)
}
```
Run dbietool with flag `-constr=factory` to generate factory objects instead of factory functions

```golang
   factory := repo.Bun[model.User]{DB: db}
   userRepo := factory.NewUser(context.Background())
```
# Naming convention

## SelectBy*|FindBy*

Can be used to select items by some criteria.

### Criteria
For now only one criteria is supported per method. 

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




### Sort order

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
