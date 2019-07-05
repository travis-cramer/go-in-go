# go-in-go
The Ancient Chinese board game Go written in Go.

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

# usage
You have the option to either run the CLI application or start the web API:
```
$ go-in-go -cli
```
or
```
$ go-in-go -api
```
The program defaults to running the CLI if no option is given.

# web
While running the API, open the `web/index.html` file in your browser to use the web interface.

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
