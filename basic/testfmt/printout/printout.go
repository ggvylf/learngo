package printout

import "fmt"

func Printout() {
	str1 := "hello world"
	str2 := `aaa
		bbb
		ccc
		`
	var str3 byte = 'c'

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Printf("%c\n", str3)

	str4 := "sdfjaksldfjalsdf"
	str5 := str4[:3]
	fmt.Println(str5)
	fmt.Printf("%T", str5)
}
