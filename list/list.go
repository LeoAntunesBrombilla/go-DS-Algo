package main

import (
	"container/list"
	"fmt"
)

func main() {
	intList := list.New()
	fmt.Println(intList)
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)
	fmt.Println(intList)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

}
