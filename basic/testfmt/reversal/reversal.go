package reversal

import "fmt"

func Reversal(str string) {
	tmpstr := []byte(str)
	strLen := len(tmpstr)
	var result []byte
	for i := 0; i < strLen; i++ {
		//	result = result + str[strLen-1-i:]
		result = append(result, tmpstr[strLen-1-i])
	}

	fmt.Printf("%s\n", result)
}
