### GoCI

A simple CI pipeline to GO code.

- Step 1: Building the program using *go build* to verify if the program structure is valid.

- Step 2: Executing tests using *go test* to ensure the program does what it's intended to do.

- Step 3: Executing *gofmt* to ensure the program's format conforms to the standards.

- Step 4: Executing *git push* to push the code to the remote shared Git repository that hosts the program code.


```
# Initialize
go mod init github.com/ebsouza/go-code/goci
```

```
# Build
go build -o goci .
```

```
# Use
./goci -p <project>
```

```
# Run All Tests
go test -v
```