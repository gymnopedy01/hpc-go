# JSON CRUD API 만들기 (Gin/Gorm)
https://www.youtube.com/watch?v=lf_kiH_NPvM

## 설치

### CompileDaemon
https://github.com/githubnemo/CompileDaemon

```shell
go get github.com/githubnemo/CompileDaemon
```

### Gin
https://gin-gonic.com/docs/quickstart/
```shell 
go get -u github.com/gin-gonic/gin
```

### Gorm
https://gorm.io/docs/
```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

```shell
CompileDaemon -command="./go-crud -my-options"
```

### GoDotEnv
https://github.com/joho/godotenv

go >= 1.17  
```shell 
go install github.com/joho/godotenv/cmd/godotenv@latest
```
go < 1.17  
```shell
go get github.com/joho/godotenv/cmd/godotenv
```
### ElephantSQL
https://www.elephantsql.com/
postgres db 생성
https://api.elephantsql.com/console/2f86e3e5-0e70-42af-9a40-a2f28015a31e/details

## Run Migrate
```shell
go run migrate/migrate.go
```

## CreatePost 
https://gorm.io/docs/create.html

## PostMan Test
https://speeding-eclipse-947341.postman.co/workspace/Go-Crud~66a892bf-7dd6-481e-a99d-a280a95855ad/request/2390863-bf684f1e-10ff-42d2-ac96-b5421c9bd653

## Query Test
https://gorm.io/docs/query.html
