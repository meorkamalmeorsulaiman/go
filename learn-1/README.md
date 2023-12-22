Learning 1

### Use golang:alpine images

`docker pull golang:alpine`

Run the `docker run -it -v ./learn-1:go --name test golang:alpine /bin/sh`

### Compiling codes

- `go run` command build temp binary and remove it one complete

```
/go # go run src/hello.go 
Hello, world!
/go # 
```

- To build the binary use `go build`

```
Hello, world!
/go # go build src/hello.go 
/go # ls
README.md  bin        docker     hello      src
/go # cd bin/
/go/bin # ls
/go/bin # cd ..
/go # ls
/go # ./hello
Hello, world!
/go # go build -o test_hello src/hello.go 
/go # ls
README.md   bin         docker      hello       src         test_hello
/go # ./test_hello 
Hello, world!
/go # 
```
- The `-o` flags allow you to name the binary and locate your src code

### Installint 3rd-party go codes

- Example:

`go install github.com/x/x`

This will download x and all it's dependencies, build and install the binary into `$PATH/bin`. That's where your binary will stored.