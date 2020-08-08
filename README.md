# go-statsd

[![Build Status][1]][2]
[![GoDoc][3]][4]
[![Go Report card][5]][6]
[![codecov][7]][8]

[1]: https://github.com/cameronnewman/go-statsd/workflows/Continuous%20Integration/badge.svg
[2]: https://github.com/cameronnewman/go-statsd/actions
[3]: https://godoc.org/github.com/cameronnewman/go-statsd?status.svg
[4]: http://godoc.org/github.com/cameronnewman/go-statsd
[5]: https://goreportcard.com/badge/github.com/cameronnewman/go-statsd
[6]: https://goreportcard.com/report/github.com/cameronnewman/go-statsd
[7]: https://codecov.io/gh/cameronnewman/go-statsd/branch/master/graph/badge.svg
[8]: https://codecov.io/gh/cameronnewman/go-statsd

## Purpose

Simple Statsd implementation

## Usage

```golang
package main

import (
    statsd "github.com/cameronnewman/go-statsd"
)

func main() {

    metrics, err := statsd.New(&statsd.Options{
        Host:      "127.0.0.1",
        Port:      9191,
        Namespace: "myapp",
    })
    if err != nil {
        panic(err)
    }

    err = metrics.Counter("Starting", 1)
    if err != nil {
        panic(err)
    }
}
```

## Issues

* None

## License

MIT Licensed
