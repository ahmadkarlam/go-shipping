# Description
This application uses Domain Driven Design, the code is divided based on existing domains.
There is only one doamin, namely the domain warehouse.
This warehouse has a feature to display a list of warehouses (warehouse code, x and y locations and current stock).
There are two endpoints in this application:
1. `/warehouse [GET]`: list warehouse
2. `/warehouse/send-vaccine [POST]`: send vaccine to patient

You can access the swagger documentation [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) to try it. 

# Project structure
1. `/cmd`
    - `/common/cli`: action that is run on the CLI.
    - `/common/http`: run http server.
2. `/common`
    - `/common/helpers`: helpers for all modules.
    - `/common/constant`: constant for all modules.
    - `/common/resolver`: struct initialization happen here.

3. `/database`: migration script.
4. `/infrastructures`: docker file.
5. `/modules`
    - `/modules/*/dto`: transform model to another object, e.g. for response purpose.
    - `/modules/*/handlers`: handle http request.
    - `/modules/*/models`: model database.
    - `/modules/*/repositories`: communicate to ORM with models.
    - `/modules/*/services`: business logic happen here.

# Database schema
In this application, only one table is used, namely:

table `warehouses`:

| Column  | Data type | Description |
|---|---|---|
| id | int |  
| code | varchar(100) | code warehouse |
| stock | int  | current stock vaccine |
| x |  int | x location |
| y | int  | y location |



# Run project
Run application with docker:
```
docker-compose up
```
Until terminal output like this:
```
go_shipping_app | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
go_shipping_app |
go_shipping_app | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
go_shipping_app |  - using env: export GIN_MODE=release
go_shipping_app |  - using code:        gin.SetMode(gin.ReleaseMode)
go_shipping_app |
go_shipping_app | [GIN-debug] GET    /warehouse                --> github.com/ahmadkarlam/go-shipping/modules/warehouses/handlers.(*WarehouseHandler).GetAll-fm (3 handlers)
go_shipping_app | [GIN-debug] POST   /warehouse/send-vaccine   --> github.com/ahmadkarlam/go-shipping/modules/warehouses/handlers.(*WarehouseHandler).SendVaccineToLocation-fm (3 handlers)
go_shipping_app | [GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
go_shipping_app | [GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
go_shipping_app | [GIN-debug] Listening and serving HTTP on :8080
```

Then you have to run migration script in another terminal tab:
```
docker-compose exec app go run cmd/cli/migrate.go
```

## Run test
```
go test ./...
```
