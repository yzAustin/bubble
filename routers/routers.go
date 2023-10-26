package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//建立gin
	r := gin.Default()
	//告诉gin框架找静态文件
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)
	//v1
	v1Group := r.Group("v1")
	{
		//待办事项

		//添加--增
		v1Group.POST("/todo", controller.CreateTodo)

		//查看--查所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//查看--查某一个待办事项

		//修改状态--改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除--删
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	//待办事项
	//添加--增
	//查看--查
	//改变状态--改
	//删除--删
	return r
}
