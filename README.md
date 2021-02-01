# gorm-test

Go Version: 1.15
Go Mod dependency manager

## Dependencies
Gorm: go get gorm.io/gorm
Gorm Postgres driver: go get gorm.io/driver/postgres
Google UUID: go get github.com/google/uuid

```
module github.com/samuelgoes/gorm-test

go 1.15

require (
	github.com/google/uuid v1.2.0 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.20.12
)
```