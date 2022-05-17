package main

import (
	"so/i"
)

type DriverA struct {
}

func (d DriverA) Name() string {
	return "hello! this is DriverA"
}

func NewDriver() i.Driver {
	return &DriverA{}
}

func main() {

}
