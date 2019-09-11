package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"study/stuMange/controller"
	"study/stuMange/data"
	"study/stuMange/repository"
	"study/stuMange/service"
)

func main() {
	app := iris.New()

	mvc.Configure(app.Party("/student"), stuHandle)

	app.Run(iris.Addr(":8080"))
}

func stuHandle(app *mvc.Application) {
	fmt.Println("初始化了")
	stuRepo := repository.NewStudentRepository(data.Students)
	stuService := service.NewStudentService(stuRepo)
	app.Register(stuService)
	app.Handle(new(controller.StudentController))
}
