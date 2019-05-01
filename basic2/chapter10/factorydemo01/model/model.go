package model

type student struct {
	Name string
	age  int
}

func NewStudent(name string, a int) *student {
	return &student{
		Name: name,
		age:  a,
	}
}

func (s *student) GetAge() int {
	return s.age
}
