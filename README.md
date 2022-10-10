# Assignment Order API (Hacktiv8 & DTS)


This API service is built using [Golang](https://go.dev/) with [gin-gonic](https://gin-gonic.com/) framework, using [GORM](https://gorm.io/) ORM with [postgreesql](https://www.postgresql.org/) database and documented with [swaggo](https://github.com/swaggo/swag)

## Instalation

1. Download this project with [git](https://git-scm.com/):
```sh
git clone https://github.com/aditgocendra/assignment_order.git
```

2. Setup your database postgreesql in root-project/database/db.go:
```sh
var (
    host     = "YOUR-HOST"
    user     = "YOUR-USERNAME-DB"
    password = "YOUR-PASSWORD-DB"
    dbPort   = "5432"
    dbName   = "orders_by"
    db       *gorm.DB
    err      error
)
```

3. Run project:
```sh
go run main.go
```
## Documentation Endpoint API

After you install, You can see the documentation of the endpoints that are automatically generated with swagger :
```sh
http://localhost:8080/swagger/index.html
```
