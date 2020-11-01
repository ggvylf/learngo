package taillog

import "fmt"

var (
	tailObj *tail.Tail
)

func Init(fileName string) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tailObj, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed,err=", err)
		return
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return <-tailObj.Lines
}
