# dbie

[![codecov](https://codecov.io/gh/iamgoroot/dbie/branch/main/graph/badge.svg?token=HDGXEOT8BA)](https://codecov.io/gh/iamgoroot/dbie)

dbie - (DB Interface Extension) Golang database layer for lazy gophers

#### Long story short:

* You [provide an interface](#define-repository-interface) of database layer using **models** from your orm
  library ([Go-pg](https://github.com/go-pg/pg), [Gorm](https://gorm.io/), [Bun](https://bun.uptrace.dev/)
  , [mongo](https://github.com/mongodb/mongo-go-driver) etc...)
* dbie **generates an implementation of that interface** for matching methods
* You can implement [custom method](#custom-methods) yourself if you need

1. [Why it might be good?](#why-it-might-be-good)
2. [Why not sqlc?](#why-not-sqlc?)
3. [What's missing](#what's-missing?)
4. [Getting started](#getting-started)
    1. [Install](#install-generator-tool)
    2. [Define contracts](#define-model)
    3. [Usage](#Usage)
5. [SelectBy*|FindBy*](#SelectBy*|FindBy*)
6. [Sort order](#sort-order)
7. [Custom methods](#custom-methods)

## Why it might be good?

* You do mostly brain-dead simple db queries ([Go-pg](https://github.com/go-pg/pg), [Gorm](https://gorm.io/)
  , [Bun](https://bun.uptrace.dev/), [mongo](https://github.com/mongodb/mongo-go-driver) etc...)
* It's a nice **addition to orm** you might already use - use pagination, sorting, filtering with
* No query pieces all over your code - **go code first**
* dbietool generates '**just enough**' code **to satisfy the interface** - for less clutter
* Generate and forget - **maintain your models** and dbie will translate the rest

## Why not sqlc?

bdie is different in a few ways so might not be for you depending on your use case or preferences

* **Go code first** approach
* **MongoDB** support
* dbie is **not a code generator**, it's a library - you don't necessarily have to generate any code (it's just your
  convenience)

## What's missing?

* No **Transactions** - but you can them implement as [custom method](#custom-methods) with your orm library
* No **Joins** - but [custom method](#custom-methods) again

#### But... I want to define interfaces at layer where I use them!

...And I encourage you to do so!

* Use interface as blueprint for your dbie implementation in database layer
* Define small interfaces in your service layer exactly where you need them (it's golang after all)

## Getting started

### Install generator tool

```sh
   go get -u github.com/iamgoroot/dbietool
   go install github.com/iamgoroot/dbietool
```

### Define repository interface

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

### Define model

As usually in Bun, Gorm, go-pg or Mongo (tag `bson`):

```golang
type User struct {
	ID       int
	Name     string
	Group    string
}
```

### Generate

That's it. generate code

   ```sh
   go generate ./...
   ```

### Usage

```golang
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

# Docs and Links

Mongo:

* https://github.com/mongodb/mongo-go-driver

Bun:

* https://bun.uptrace.dev/guide/golang-orm.html
* https://github.com/uptrace/bun

Gorm:

* https://gorm.io/docs/
* https://github.com/go-gorm/gorm

go-pg:

* https://github.com/go-pg/pg
* https://pg.uptrace.dev/
