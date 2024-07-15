package main

import "fmt"

type ObjectState int

const (
	Created = iota + 1
	InProgess
	Done
)

var stateName = map[ObjectState]string{
	Created:   "created",
	InProgess: "inProgress",
	Done:      "done",
}

func (os ObjectState) String() string {
	return stateName[os]
}

func main() {
	var state ObjectState = 1
	fmt.Printf("Numerical value: %d \n", state)
	fmt.Println("String value: ", state.String())
}
