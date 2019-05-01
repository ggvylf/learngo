package testscanf

import (
	"fmt"
)

func TestScanf() {
	var a int
	var b string
	fmt.Println("start input")
	fmt.Scanf("%d%s", &a, &b)
	fmt.Printf("a is %d,b is %s\n", a, b)

}
