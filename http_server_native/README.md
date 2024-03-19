### 1 - Starting this project
```
go mod init example/http-server
go run main.go
```


#### POST

curl -X POST localhost:8080/notes -H 'Content-Type: application/json' -d '{"id": 1, "title": "AA", "message": "BB"}' 
curl -X POST localhost:8080/notes -H 'Content-Type: application/json' -d '{"id": 2, "title": "CC", "message": "BB"}' 

#### GET

curl localhost:8080/notes

#### GET {id}

curl localhost:8080/notes/1

#### DELETE {id}

curl -X DELETE localhost:8080/notes/1