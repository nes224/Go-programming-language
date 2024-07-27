package main

import "fmt"

func act(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Value %v is %T\n", v, v)
	case string:
		fmt.Printf("Value %v is %T\n", v, v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	buildType := [...]interface{}{1,"Hello", false}
	for i := range buildType {
		act(buildType[i])
	}
}