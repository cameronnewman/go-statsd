# go-statsd

[![GoDoc](https://godoc.org/github.com/cameronnewman/lib.statsd?status.svg)](http://godoc.org/github.com/cameronnewman/lib.statsd) 
[![Report card](https://goreportcard.com/badge/github.com/cameronnewman/lib.statsd)](https://goreportcard.com/report/github.com/cameronnewman/lib.statsd)
![Build Status](https://github.com/cameronnewman/go-statsd/workflows/Continuous%20Integration/badge.svg)

## Purpose ##

Simple Statsd implementation

## Usage

```
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
