package main

import "fmt"

//student
type Student struct {
	name          string
	id            int
	student_score Score
}

//score
type Score struct {
	subject_name  string
	teacher_name  string
	subject_score int
}

// only read
func (student Student) View_Score() {
	fmt.Printf("%s의 %s 과목 성적은 %d입니다.\n", student.name, student.student_score.subject_name, student.student_score.subject_score)
}

// write
func (student *Student) Input_Score(subject_name string, teacher_name string, subject_score int) {
	student.student_score.subject_name = subject_name
	student.student_score.teacher_name = teacher_name
	student.student_score.subject_score = subject_score
}

func main() {
	fmt.Println("첫번째 학생")
	person1 := Student{name: "Hippo", id: 10000}
	person1.Input_Score("Math", "Pythagoras", 95)
	person1.View_Score()
	person1.Input_Score("Korean", "King SaeJong", 100)
	person1.View_Score()

	fmt.Println("\n두번째 학생")
	person2 := Student{name: "Mans", id: 1}
	person2.Input_Score("History", "Kant", 88)
	person2.View_Score()
	person2.Input_Score("Computer", "Dennis Ritchie", 91)
	person2.View_Score()
}
