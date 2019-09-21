package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hprose/hprose-golang/rpc"
	"net/http"
)

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello) //发布服务
	r := gin.Default()
	r.Any("/path", func(c *gin.Context) { //RPC调用
		service.ServeHTTP(c.Writer, c.Request)
	})
	r.POST("test", ginTest)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8079") // listen and serve on 0.0.0.0:8080
}

type DATA struct {
	Id string `json:"id"`
	Callback string `json:"callback"`
}



func init()  {
	fmt.Println("服务启动")
}

func hello(name string) string {
	return "Hello " + name + "!"
}

func ginTest(c *gin.Context) {


	//ids := c.QueryMap("ids") //GET
	//names := c.PostFormMap("names") //POST
	//fmt.Printf("ids: %v; names: %v", ids, names)
	//return
	//m := c.PostFormMap("test")
	//for k, v := range m {
	//	fmt.Println(k, v)
	//	//fmt.Println(v)
	//}

	var test_data DATA
	data,_ := c.GetRawData()
	json.Unmarshal(data, &test_data)
	fmt.Printf("%s", data)
	fmt.Println(test_data.Id)

	c.JSON(http.StatusOK, gin.H{
		"message" : "pomg223",
	})
}