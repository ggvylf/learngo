package main

import (
	"errors"
	"fmt"
	"time"
)

// 可复用对象
type ReUseObj struct{}

// 对象池 用来存放对象的channel
type ObjPool struct {
	buffChan chan *ReUseObj
}

// 创建对象池，入参是复用对象的个数
func NewObjPool(num int) *ObjPool {
	// 创建对象池实例
	objPool := ObjPool{}
	objPool.buffChan = make(chan *ReUseObj, num)

	// 填充对象
	for i := 0; i < num; i++ {
		objPool.buffChan <- &ReUseObj{}

	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReUseObj, error) {
	select {
	case ret := <-p.buffChan:
		return ret, nil

	// 超时判断
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReUseObj) error {
	select {
	// 把复用对象放回channel
	case p.buffChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}

}

func main() {
	pool := NewObjPool(10)

	// 正常取放
	for i := 0; i < 11; i++ {

		// 从channel里取obj
		if v, err := pool.GetObj(1 * time.Second); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%T\n", v)

			// 放回channel
			if err := pool.ReleaseObj(v); err != nil {
				fmt.Println(err)
			}
		}

	}

	// 向channel里放入多余的Obj，会报overflow
	if err := pool.ReleaseObj(&ReUseObj{}); err != nil {
		fmt.Println(err)
	}

	// 只取不放，最后一个拿不到会报timeout
	for i := 0; i < 11; i++ {

		if v, err := pool.GetObj(1 * time.Second); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%T\n", v)
		}

	}

	fmt.Println("Done")

}
