# go-in-go
A CLI version of the Ancient Chinese board game Go written in Golang

# install and run
The following assumes that you have your GOBIN setup and added to your PATH (see: https://golang.org/doc/code.html).

Get all dependencies and any dependencies of dependencies:
```
$ go get ./...
```

Install the binary to your GOBIN:
```
$ go install
```
Run the binary:
```
$ go-in-go
```


# testing
In the root directory of the repository, run

```
$ go test
```

to run our test suite. See: https://golang.org/pkg/testing/
