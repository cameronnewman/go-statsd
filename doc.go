/*

Package statsd is simple statsd warpper.

*/
package statsd

// Simple example:
//
// package main
//
//import (
//	statsd "github.com/cameronnewman/lib.statsd"
//)
//
//func main() {
//
//	metrics, err := statsd.New(&statsd.Options{
//		Host:      "127.0.0.1",
//		Port:      9191,
//		Namespace: "myapp",
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	err = metrics.Counter("Starting", 1)
//	if err != nil {
//		panic(err)
//	}
//}
//
