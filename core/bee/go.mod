module github.com/iamgoroot/dbie/core/bee

go 1.18

require (
	github.com/beego/beego/v2 v2.0.4
	github.com/iamgoroot/dbie v0.0.0-20220715232405-9fcd7d479f61
)

replace github.com/iamgoroot/dbie => ../../

require (
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
)
