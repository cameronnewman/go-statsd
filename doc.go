/*

Package statsd is simple statsd warpper lib

*/
package statsd

// Simple example:
//
// package main
//
// import (
// 	"fmt"
// 	"github.com/cameronnewman/lib.statsd"
// )
//
// func main() {
//
//	sd, err := statsd.New(&statsd.Config{
//		StatsdHost: "127.0.0.1",
//		StatsdPort: 9191,
//		StatsdNamespace: "myapp"
//	})
//
//	if err != nil{
//		panic(err)
//	}
//
//	sd.StatsCounter("Starting", 1)
//
// }
