package main

import "fmt"

func main() {
	var scores [3][5]float64

	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班第%d个学生的成绩:", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	totalsum := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		fmt.Printf("第%d个班级的总分=%f,平均分=%f\n", i, sum, sum/float64(len(scores[i])))
		totalsum += sum
	}

	fmt.Printf("所有班级的的总分=%f,平均分为=%f\n", totalsum, totalsum/float64(len(scores)*len(scores[0])))

}
