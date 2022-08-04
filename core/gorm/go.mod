module github.com/iamgoroot/dbie/core/gorm

go 1.18

require (
	github.com/iamgoroot/dbie v0.0.0-20220715232405-9fcd7d479f61
	gorm.io/gorm v1.23.8

)

replace github.com/iamgoroot/dbie => ../../

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)
