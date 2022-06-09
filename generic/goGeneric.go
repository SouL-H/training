package main

import "fmt"

func main() {
	genericType("hello")
	genericType(1)
	genericType(true)
	genericType([]int{1, 2, 3})
	genericType(map[string]int{"a": 1, "b": 2})
	genericTypeV2("hello")
	genericTypeV2(1)
	genericTypeV2(true)
	genericTypeV2([]int{1, 2, 3})
	genericTypeV2(map[string]int{"a": 1, "b": 2})

}

func genericType[T any](say T) {
	fmt.Println(say)
}
func genericTypeV2(say interface{}) {
	fmt.Println(say)
}
