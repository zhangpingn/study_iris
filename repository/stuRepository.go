package repository

import (
	"fmt"
	"study/stuMange/model"
	"sync"
)

type StudentRepository struct {
	lock sync.Mutex
	source map[string]*model.Student
}

func NewStudentRepository(source map[string]*model.Student) StudentRepository {
	return StudentRepository{
		source:source,
	}
}

func (this *StudentRepository) GetStudentBySid(sid string) *model.Student {
	this.lock.Lock()
	defer this.lock.Unlock()
	var(
		student *model.Student
		ok      bool
	)
	if student, ok = this.source[sid]; ok {
		fmt.Printf("find %s: %v\n", sid, student)
		return student
	}

	return nil
}

func (this *StudentRepository) AddOrUpdateStudent(student *model.Student) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	if len(this.source) == 0 {
		this.source = make(map[string]*model.Student, 1)
	}
	this.source[student.Sid] = student
	return true
}

func (this *StudentRepository) DeleteStudentBySid(sid string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.source, sid)
	return true
}

func (this *StudentRepository) GetAll() map[string]*model.Student {
	return this.source
}