package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := `
<div>hello world</div>
<div>heheh</div>
<div>xuxy</div>
<div>
中文测试
中文
</div>
`
	//单行模式，?s是模式修饰符，表示更改的含义,可以修改去匹配每个字符，包括\n,。表示单个字符，*?表示匹配任意次
	ret := regexp.MustCompile(`<div>(?s:(.*?))</div>`)


	alls := ret.FindAllStringSubmatch(str, -1)



	//每个v的返回值都有2个，一个是带匹配参考项的，一个不包含
	for index, one := range alls {
		fmt.Printf("the %d key  first value = %v,the  second value =%v\n", index+1, one[0], one[1])
	}

}
