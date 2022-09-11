package main

import (
	"context"
	"github.com/go-pg/pg/extra/pgdebug/v10"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/iamgoroot/dbie/core/test/model"
	"github.com/iamgoroot/dbie/core/test/repo"
	"log"
)

func main() {
	// instantiate using factory
	factory := repo.Pg{DB: connect()}
	userRepo := factory.NewUser(context.Background())

	// insert user
	err := userRepo.Insert(model.User{Name: "userName1"})
	if err != nil {
		log.Fatalln(err)
	}

	// select user using generated method
	user, err := userRepo.SelectByName("userName1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(user, err)
}

func connect() *pg.DB {
	dsn := ""
	opt, err := pg.ParseURL(dsn)
	if err != nil {
		log.Fatalln(err)
	}
	db := pg.Connect(opt)
	db.AddQueryHook(pgdebug.NewDebugHook())
	db.Model(&model.User{}).Context(context.Background()).DropTable(&orm.DropTableOptions{Cascade: true})
	err = db.Model(&model.User{}).Context(context.Background()).CreateTable(&orm.CreateTableOptions{FKConstraints: true})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
