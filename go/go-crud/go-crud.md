# [Creating a JSON CRUD API in Go (Gin/GORM)](https://www.youtube.com/watch?v=lf_kiH_NPvM)
JSON CRUD API 만들기 (Gin/Gorm) - what.BlogPost  

## 설치

### [CompileDaemon](https://github.com/githubnemo/CompileDaemon) - go 파일 자동 컴파일및 실행
```shell
go get github.com/githubnemo/CompileDaemon
```

### [Gin](https://gin-gonic.com/docs/quickstart/) - http web framework
```shell 
go get -u github.com/gin-gonic/gin
```

### [Gorm](https://gorm.io/docs/) - ORM Library
```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

```shell
CompileDaemon -command="./go-crud -my-options"
```

### [GoDotEnv](https://github.com/joho/godotenv)
- go >= 1.17  
```shell 
go install github.com/joho/godotenv/cmd/godotenv@latest
```
- go < 1.17  
```shell
go get github.com/joho/godotenv/cmd/godotenv
```

### [ElephantSQL](https://www.elephantsql.com/) PostgresSQL Saas
- postgres db 생성
  - https://api.elephantsql.com/console/2f86e3e5-0e70-42af-9a40-a2f28015a31e/details

### Run Migrate
[migrate.go](migrate/migrate.go) Table 초기 생성 및 구동법
```shell
go run migrate/migrate.go
```

### CreatePost 
docs : https://gorm.io/docs/create.html  
[main.go](main.go)  
[poostController](controllers/postController.go)

### IndexPost
docs : https://gorm.io/docs/query.html

### GetPost 
docs : https://gorm.io/docs/query.html

### UpdatePost
docs : https://gorm.io/docs/query.html

### DeletePost
docs : https://gorm.io/docs/query.html

### PostMan Test
[go-crud.http](go-crud.http)  
[PostMan 생성](https://speeding-eclipse-947341.postman.co/workspace/Go-Crud~66a892bf-7dd6-481e-a99d-a280a95855ad/request/2390863-bf684f1e-10ff-42d2-ac96-b5421c9bd653)
