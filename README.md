# go-in-go
A CLI version of the Ancient Chinese board game Go written in Golang

# install and run
The following assumes that you have the basics of your go workspace setup (see: https://golang.org/doc/code.html).

Get all dependencies and any dependencies of dependencies (as well as dependencies for tests):
```
$ go get -t ./...
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

# todo
Things that we still want to get done
* Implement how to end the game and how to calculate the winner
* Building it out into a web app and browser game
    * Hovering over a square shows a greyed out piece
* Use a redis instance to remember game state
* Dockerize it all
