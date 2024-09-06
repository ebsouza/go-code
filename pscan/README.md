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
