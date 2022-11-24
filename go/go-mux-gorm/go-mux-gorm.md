# [Creating Golang Rest API with Mux and GORM - Complete Tutorial](https://www.youtube.com/watch?v=KPftgI40WHI)

## Install

``` shell
go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## Run Server
```shell
go build
./go-mux-gorm.exe 
./
```

## Development
### Router setting
```go
func main () {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
}
```

### Func 
#### - Youtube style
```go
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}
```
#### Document Style 
```go
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

```