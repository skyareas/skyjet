module github.com/skyareas/skyjet/examples/session

go 1.15

require (
	github.com/skyareas/skyjet v0.14.0
	gorm.io/driver/sqlite v1.3.1 // indirect
	gorm.io/gorm v1.23.2 // indirect
)

replace github.com/skyareas/skyjet v0.14.0 => ../..
