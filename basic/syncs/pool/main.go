package main

import (
	"fmt"
	"sync"
)

// sync.Pool对象的获取
// 从私有对象里获取，这是协程安全的
// 私有对象里没有，从Processor的共享池里获取，是协程不安全的
// 当前Processor的共享池是空的，去其他Processor的共享池里获取
// 如果所有的共享池都是空的 用户用New函数产生一个新的对象再返回

//sync.Pool对象的放回
// 私有对象不存在则保存为私有对象，这是协程安全的
// 当私有对象存在的时候，放入当前Processor的共享池，这是协程不安全的

//sync.Pool的生命周期
// GC会清楚sync.Pool中缓存的对象
// 缓存对象的有效期是在下一次GC之前

//sync.Pool的使用场景
//1. 适合于通过复用，降低复杂对象创建和GC代价
//2. 协程安全，但是有锁的开销
//3. 生命周期收到GC影响，不适合做连接池等需要自己管理生命周期的资源

func getAndPut() {
	//声明pool的时候已经创建了一个100的对象
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new pool object")
			return 100
		},
	}

	//接口类型需要做断言
	//取出pool的对象
	v := pool.Get().(int)
	fmt.Println(v)

	//把3放入pool中
	//上次已经取了100，pool是空的，3是会被当成私有对象处理
	pool.Put(3)

	//可以强制gc来看效果,私有对象的有效期在下次GC之前
	//GC以后，私有对象将被删除，获取的对象是实例化的100
	// runtime.GC()

	v1, _ := pool.Get().(int)
	fmt.Println(v1)

}

// 多个goroutine下的sync.Pool
// 开了10个goutine来从pool中取对象
// 先把3个10取出来，然后pool是空的
// 再取的时候会调用New函数生成的100
func syncPoolInMultiGoruntine() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new pool object")
			return 100
		},
	}

	pool.Put(10)
	pool.Put(10)
	pool.Put(10)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get().(int))
			wg.Done()
		}(i)

	}
	wg.Wait()

}

func main() {
	// getAndPut()
	syncPoolInMultiGoruntine()

}
