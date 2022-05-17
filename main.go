package main

import (
	"fmt"
	"plugin"

	"so/i"
)

func main() {
	pl, err := plugin.Open("./hello/hello.so")
	if err != nil {
		panic(err)
	}
	newDriverSymbol, err := pl.Lookup("NewDriver")
	if err != nil {
		panic(err)
	}

	newDriverFunc := newDriverSymbol.(func() i.Driver)
	newDriver := newDriverFunc()
	fmt.Println(newDriver.Name())
}
