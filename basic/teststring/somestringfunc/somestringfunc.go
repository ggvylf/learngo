package main

import (
	"fmt"
	"strings"
)

//替换str中old为new，count表示替换的个数，为-1表示不限制
func testReplace(str, old, new string, count int) string {
	result := strings.Replace(str, old, new, count)
	return result
}

// 计算old在str中的出现的个数
func testCount(str, old string) int {
	result := strings.Count(str, old)
	return result
}

//重复count次输出str
func testRepeat(str string, count int) string {
	result := strings.Repeat(str, count)
	return result
}

func main() {
	var (
		str   string
		old   string
		new   string
		count int
	)
	fmt.Scanf("%s%s%s%d", &str, &old, &new, &count)

	replaced := testReplace(str, old, new, count)
	fmt.Println(replaced)

	counter := testCount(str, old)
	fmt.Println(counter)

	repeater := testRepeat(str, count)
	fmt.Println(repeater)

	//小写转大写
	REP := strings.ToUpper(repeater)
	fmt.Println(REP)

	//大写转小写
	fmt.Println(strings.ToLower(REP))

}
