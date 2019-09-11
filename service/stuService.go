package service

import (
	"study/stuMange/model"
	"study/stuMange/repository"
)

type StudentService struct {
	studentRepository repository.StudentRepository
}

func NewStudentService(studentRepository repository.StudentRepository) StudentService{
	return StudentService{
		studentRepository:studentRepository,
	}
}

func (this *StudentService) GetStudentBySid(sid string) *model.Student {
	return this.studentRepository.GetStudentBySid(sid)
}

func (this *StudentService) GetAll() map[string]*model.Student {
	return this.studentRepository.GetAll()
}

func (this *StudentService) DeleteStudentBySid(sid string) bool {
	return this.studentRepository.DeleteStudentBySid(sid)
}

func (this *StudentService) AddOrUpdateStudent(student *model.Student) bool {
	return this.studentRepository.AddOrUpdateStudent(student)
}