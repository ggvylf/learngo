package main

import "fmt"

type Student struct {
	Name  string
	Score float64
}

func (s *Student) ShowScore() {
	fmt.Println("scores=", s.Score)
}

func (s *Student) SetScore(score float64) {
	s.Score = score

}

func (s *Student) GetSum(m, n int) int {
	return m + n
}

type Pupil struct {
	Student
}

func (p *Pupil) Testing() {
	fmt.Println("is examing")
}

type Graduate struct {
	Student
}

func (g Graduate) Testing() {
	fmt.Println("is examing")
}

func main() {
	p := &Pupil{}
	p.Student.Name = "jerry"

	p.Student.SetScore(500)

	fmt.Println(p)

	fmt.Println(p.Student.GetSum(2, 3))

	g := &Graduate{}

	g.Student.SetScore(300)

	fmt.Println(g)
	g.Student.ShowScore()

}
