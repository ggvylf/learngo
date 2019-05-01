package main

import "fmt"

var container = []string{"one", "two", "three"}

func main() {
	container := map[int]string{0: "one", 1: "two", 2: "three"}

	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)

	if !(ok1 || ok2) {
		fmt.Printf("not suppport :%T\n", container)
		return
	}
	fmt.Printf("element is %q,type is %T\n", container[1], container)

	elem, err := getElement(container)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("element is %q,type is %T\n", elem, container)

}
func getElement(containerI interface{}) (elem string, err error) {

	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("not suppport :%T\n", container)
		return
	}
	return
}
