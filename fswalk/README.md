### File System Walker

```
go mod init github.com/ebsouza/go-code/fswalk
```

#### Building the application
```
go build -o fswalk main.go actions.go
```

#### Running application
```
./fswalk -root <dir> -del -archive <dir> -ext <ext> -size <size> -logFile <file>
```


#### Running unit tests
```
go test -v
```