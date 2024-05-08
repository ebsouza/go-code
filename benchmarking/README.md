## Benchmarking

Golang includes built-in tools for writing benchmarks in the testing package and the go tool, so you can write useful benchmarks without installing any dependencies.


### Running benchmarking tests
```
go test -bench=.
```

There are some important flags
- **count**: sets the number of times each test will run;
    - Example: go test -bench=. -count=5
- **benchtime**("Ns"): sets the minimum amount of time that the benchmark function will run;
    - Example: go test -bench=. -benchtime=2x
- **benchtime**("Nx"): sets the number of iterations of each test;
    - Example: go test -bench=. -benchtime=10x
- **benchmem**: include memory allocation statistics.
    - Example: go test -bench=. -benchmem
- **-run=^#**: skip all unit tests
    - Example: go test -bench=. -run=^#