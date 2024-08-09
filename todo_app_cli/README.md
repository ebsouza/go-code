### Simple ToDo CLI App

```
go mod init github.com/ebsouza/todo_app_cli
```

```
go build -o todo_app cmd/cli/main.go
```

#### Commands

```
./todo_app -task Task1
```

```
./todo_app -add Task1
```

```
echo "Task1" ./todo_app -add
```

```
./todo_app -list
```

```
./todo_app -complete <Task Index>
```

```
./todo_app -h
```