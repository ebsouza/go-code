### pScan


#### Setup environment
```
go install github.com/spf13/cobra-cli@latest
```

```
# add cobra-cli to PATH
export PATH=$(go env GOPATH)/bin:$PATH
```


#### Init application
```
go mod init github.com/ebsouza/go-code/pscan
cobra-cli init
```

#### Run tests
```
go test ./...
```

#### Build application
```
go build -o pScan main.go
```

#### Run application
```
./pScan hosts add localhost
./pScan hosts list
./pScan scan --ports 22,80,443
```


#### Generate Documentation
```
mkdir docs
./pScan docs --dir ./docs
```
