package main

import "fmt"

//引用传递保证地图不变，i j 表示要探索的点的坐标，返回探测的结果
func SetWay(myMap *[8][7]int, i int, j int) bool {

	//迷宫右下角就是出口
	if myMap[6][5] == 2 {
		return true
	} else {
		//没有探测过的点，可以探索
		if myMap[i][j] == 0 {
			//假设当前探测的点是通的,探索的的策略为下右上左，探索不通就换方向

			myMap[i][j] = 2

			//向下探索
			if SetWay(myMap, i+1, j) {
				return true
				//向右探索
			} else if SetWay(myMap, i, j+1) {
				return true
				//向上探索
			} else if SetWay(myMap, i-1, j) {
				return true
				//向左探索
			} else if SetWay(myMap, i, j-1) {
				return true
				//都不通，是死路
			} else {
				//原先的假设是错误的，标记为3
				myMap[i][j] = 3
				return false
			}

		} else {
			//该点不能探测，是墙
			return false
		}

	}

}

func main() {
	//定义一个迷宫，数组元素的值=1表示墙，元素的值=0表示还没有探索过的路径，元素的值=2表示通路，元素的值=3表示已经走过，但是是死路
	var myMap [8][7]int

	//定义围墙
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1

	}

	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1

	}

	//迷宫内加墙
	myMap[3][1] = 1
	myMap[3][2] = 1

	//打印地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	SetWay(&myMap, 1, 1)

	//打印地图
	fmt.Println("探测完毕")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

}
