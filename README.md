# Golang Rest Full Shop API (Gin Freamwork)

## This Project Feature
- Using Gin Freamwork 
- Clean Architecture
- Migrator Package In Codebase 
- Migrations With CLI Tool
- Advance Project Configurations
- Using Docker For Setup Project Dependencies 


### CLI Tool For Migration
migration up:
```
migrate -source file://migrations -database "mysql://shop:shop-pass@(localhost:3308)/shop?parseTime=true&x-migrations-table=migrations" up 
```
migration down:
```
migrate -source file://migrations -database "mysql://shop:shop-pass@(localhost:3308)/shop?parseTime=true&x-migrations-table=migrations" down 
```