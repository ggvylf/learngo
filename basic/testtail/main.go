package main

import (
	"fmt"
	"time"

	//第三方包
	"github.com/hpcloud/tail"
)

func main() {
	//配置
	config := tail.Config{
		//等同于tail -F或者tailf
		Follow: true,
		//自动重新打开
		ReOpen: true,
		//文件是否必须存在
		MustExist: false,
		//轮循文件更改，而不是使用inofify
		Poll: true,
		//读取的偏移量，和读取到哪里，参考os.SEEK_*，这里设置为读取到结尾
		//Location: &tail.SeekInfo{Offset: 0, Whence: 2},
	}

	//要监听的文件
	filename := "./test.log"
	t, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Println(err)
	}

	//遍历tail到的结果
	for {
		//从管道中取出数据
		msg, ok := <-t.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", t.Filename)
			time.Sleep(100 * time.Microsecond)
			continue

		}
		fmt.Println("msg=", msg)
	}

}
