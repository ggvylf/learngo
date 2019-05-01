package main

import "fmt"

//结构体
type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//遍历原始数组
	fmt.Println("原始数组：")
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
	fmt.Println("------------------")

	//稀疏数组的方法
	//1.声明一个ValNode的切片
	//2. 遍历原始数组，如果发现有元素不为0的，放入切片中

	var spareArr []ValNode

	//创建稀疏数组头，包含行、列、默认值
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}

	spareArr = append(spareArr, valNode)

	for row, v := range chessMap {
		for col, val := range v {
			if val != 0 {
				valNode := ValNode{
					row: row,
					col: col,
					val: val,
				}
				spareArr = append(spareArr, valNode)
			}
		}
	}

	fmt.Println("稀疏数组：")
	//输出稀疏数组
	for i, valNode := range spareArr {
		fmt.Printf("index=%d: row=%d col=%d value=%d\n", i, valNode.row, valNode.col, valNode.val)

	}
	fmt.Println("------------------")

	//恢复成原始数组

	var chessMap2 [11][11]int

	for _, valNode := range spareArr[1:] {
		chessMap2[valNode.row][valNode.col] = valNode.val
	}

	fmt.Println("恢复后的数组：")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
	fmt.Println("------------------")

}
