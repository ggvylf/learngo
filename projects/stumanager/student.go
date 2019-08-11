package main

import "fmt"

type student struct {
	id    int    `json:"id"`
	name  string `json:"name"`
	class string `json:"class"`
}

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudent []*student
}

func newStudentMgr(cap int) *studentMgr {
	return &studentMgr{
		allStudent: make([]*student, 0, cap),
	}
}

//添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudent = append(s.allStudent, newStu)

}

//编辑学生
func (s *studentMgr) editStudent(newStu *student) {
	for i, v := range s.allStudent {
		if newStu.id == v.id {
			s.allStudent[i] = newStu
			fmt.Printf("学号：%d,姓名:%s,班级：%s\n", v.id, v.name, v.class)
		}
	}
	fmt.Println("学生不存在")
}

//展示学生
func (s *studentMgr) showallStudent() {
	for _, v := range s.allStudent {
		fmt.Printf("学号：%d,姓名:%s,班级：%s\n", v.id, v.name, v.class)
	}
}
