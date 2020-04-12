//中间件的主要作用是再处理请求的过程中，允许加入用户自定义的hook函数。
//中间件适合处理一些公共的业务逻辑。例如登录 权限校验 数据分页，记录日志，耗时等。
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func myindexHandler(c *gin.Context) {
	fmt.Println("myindexHandler start")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
	fmt.Println("myindexHandler stop")
}

//定义一个middleware m1
func m1(c *gin.Context) {
	fmt.Println("middleware m1 start")
	//执行后续的处理函数
	c.Next()
	//阻止后续的处理函数的调用
	//c.Abort()s
	fmt.Println("middleware m1 stop")
}

//定义一个middleware m2
func m2(c *gin.Context) {
	fmt.Println("middleware m2 start")
	c.Next()
	fmt.Println("middleware m2 stop")
}

//中间件的类型必须是gin.HandlerFunc
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		//在context中制定值
		c.Set("name", "tom")
		fmt.Println("statcost start")
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		fmt.Printf("cost=%v\n", cost)
		fmt.Println("statcost stop")
	}
}

func myAuth(run bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if run {
			fmt.Println("start auth")
			fmt.Println("stop auth")
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func mygoroutine() gin.HandlerFunc {
	return func(c *gin.Context) {
		//goroutine中只能使用context的拷贝，不能直接修改context的值
		myc := c.Copy()
		go func() {
			name := myc.MustGet("name").(string)
			fmt.Println("myc.name=", name)
		}()

	}
}

func main() {

	r := gin.Default()

	//默认情况下，gin.Default()使用了Logger和Recovery中间件
	// 定义没有默认中间件的路由
	// r:=gin.New()

	//全局注册中间件，注册后不需要明确指定
	// r.Use(StatCost())

	r.GET("/index", StatCost(), myindexHandler)

	//多个中间件的执行顺序，先执行myindex 然后m2. 后执行m1
	r.GET("/multimiddleware", m1, m2, myindexHandler)

	//给中间件增加判断开关
	r.GET("/auth", myAuth(false), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "/auth",
		})

	})

	//route group使用中间件
	//方法1
	myGroup1 := r.Group("/mygroup1", StatCost())
	{
		myGroup1.GET("/a", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "mygroup1/a",
			})
		})
	}

	//方法2
	myGroup2 := r.Group("/mygroup2")
	{
		myGroup2.Use(StatCost())
		myGroup2.GET("/a", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "mygroup2/a",
			})
		})
	}

	r.GET("/testgoroutine", StatCost(), mygoroutine())

	r.Run(":8888")
}
