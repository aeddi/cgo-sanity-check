package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import "../test.h"
*/
import "C"

func main() {
	C.printHelloWorld()
}
