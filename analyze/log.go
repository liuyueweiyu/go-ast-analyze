package main

import (
	"fmt"
)

// Log Log
func Log(args ...interface{}) {
	fmt.Println(args...)
}

// SLog SLog
func SLog(args ...interface{}) {
	fmt.Printf("%+v\n", args)
}
