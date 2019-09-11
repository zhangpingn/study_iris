package controller

import (
	"study/stuMange/model"
	"study/stuMange/service"
)

type StudentController struct {
	studentService service.StudentService
}

func (this *StudentController) GetBy(sid string) *model.Student {
	return this.studentService.GetStudentBySid(sid)
}