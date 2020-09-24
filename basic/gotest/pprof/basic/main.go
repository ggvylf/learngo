package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Println("recv from.chan,value=%v\n", v)
		default:
		}
	}
}

func main() {

	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")

	flag.Parse()

	if isCPUPprof {
		f1, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Println("create cpu pporf file failed,err=", err)
			return
		}

		pprof.StartCPUProfile(f1)
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()

	}

	for i := 0; i < 6; i++ {
		go logicCode()
	}

	time.Sleep(20 * time.Second)

	if isMemPprof {
		f2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Println("create mem pporf file failed,err=", err)
			return
		}

		pprof.WriteHeapProfile(f2)
		f2.Close()

	}

}
