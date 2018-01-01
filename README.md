# lib.statsd
Simple Statsd implementation

## Dev guide

### 1. [Install Go](https://golang.org/doc/install)

### 2. Set your $GOPATH variable, [read more here](https://github.com/golang/go/wiki/SettingGOPATH)

$GOPATH should point to a directory where you store ( or intend to store) all of your go projects.

```sh
export GOPATH=$HOME/projects/all-the-go-projects
```

### 3. Install govendor.

We are using [`govendor`](https://github.com/kardianos/govendor). Install it:

```sh
go get -u github.com/kardianos/govendor/...
```

To get a dependency into a vendor directory:

```sh
govendor fetch github.com/cameronnewman/lib.statsd/
```

To synch dependencies with the vendor directory:

```sh
govendor update +external
```

### 5. Clone the repo with `go get` command

```sh
cd $GOPATH
go get github.com/cameronnewman/lib.statsd/...
```

## Tests

```sh
make test
```
