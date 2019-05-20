package main

import (
	"fmt"
	"strings"
)

func main() {
	// 示例1。
	var builder1 strings.Builder
	builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	fmt.Printf("The first output:\nlen:%d\nstrings:%q\ncap:%d\n", builder1.Len(), builder1.String(), builder1.Cap())
	fmt.Println()
	builder1.WriteByte(' ')
	builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
	builder1.Write([]byte{'\n', '\n'})
	builder1.WriteString("Do not copy a non-zero Builder.")
	fmt.Printf("The first output:\nlen:%d\nstrings:%q\ncap:%d\n", builder1.Len(), builder1.String(), builder1.Cap())
	fmt.Println()

	// 示例2。
	fmt.Printf("before grow:\nlen:%d\ncap:%d\n", builder1.Len(), builder1.Cap())
	fmt.Println("Grow the builder ...")
	builder1.Grow(10)
	//len没有超过cap，所以Grow没有生效
	fmt.Printf("after grow:\nlen:%d\ncap:%d\n", builder1.Len(), builder1.Cap())
	fmt.Println()

	// 示例3。
	fmt.Println("Reset the builder ...")
	builder1.Reset()
	fmt.Printf("After Reset output:\nlen:%d\nstrings:%q\ncap:%d\n", builder1.Len(), builder1.String(), builder1.Cap())
}
