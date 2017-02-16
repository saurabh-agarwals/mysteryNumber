## Setup tests.

```
	Install Go version >= 1.6
	See https://golang.org/doc/install or doc/install.html for more details.

	mkdir mysteryNumber
	cd mysteryNumber
	export GOPATH=`pwd`
	export GOBIN=$GOPATH/bin
```

## Fetch code
```
	go get github.com/saurabh-agarwals/mysteryNumber/grab  
```

## Steps to Start Tests
```
	go test -v $(go list ./... | grep -v /vendor/)
```