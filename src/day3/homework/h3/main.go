package main


import (
	"fmt"
	// "encoding/json"
	

)


type Student struct {
	Username string `json:"username"`
	Grade int `json:"grade"`
	Gender string `json:"gender"`
	Score int `json:"score"`
}

var student []Student = make([]Student,0,0)

var flag bool = false

func AddStuInfo() {
	var stu Student
	fmt.Println("Please input student name:")
	fmt.Scanf("%s\n",&stu.Username)
	fmt.Println("Please input student grade:")
	fmt.Scanf("%d\n",&stu.Grade)
	fmt.Println("Please input student gender:")
	fmt.Scanf("%s\n",&stu.Gender)
	fmt.Println("Please input student score:")
	fmt.Scanf("%d\n",&stu.Score)
    fmt.Print("Name\tGrade\tGender\tScore\n")
	fmt.Printf("%v\t%v\t%v\t%v\n",stu.Username,stu.Grade,stu.Gender,stu.Score)
	student = append(student,stu)
}

func ModifyUserInfo() {

	fmt.Printf("Please input the name which you want to modify:")

	var studentName string
	fmt.Scanf("%v\n", &studentName)
	for k, stu := range student {
		if stu.Username == studentName {

			var name string
			fmt.Printf("Please input the new name:\n")
			fmt.Scanf("%s\n", &name)
			stu.Username = name
			fmt.Printf("Please input the new score\n")
			var score int
			fmt.Scanf("%v\n", &score)
			stu.Score = score

			fmt.Printf("Please input the new grade\n")
			var grade int
			fmt.Scanf("%v\n", &grade)
			stu.Grade = grade

			fmt.Printf("Please input the new gender\n")
			var gender string
			fmt.Scanf("%v\n", &gender)
			stu.Gender = gender
			student[k] = stu
			break
		}

	}
	fmt.Printf("The name is not exist!\n")

}

func run() {

	var selectNum int

	fmt.Print("\t\t\t\tWelcome to the student manage system\t\t\t\n")
	fmt.Print("#######################################################################################################################\n")
	fmt.Print("Please select the NO:\n")
	fmt.Print("1)Add new student information\n2)Modify student information\n3)Display all student information\n4)Quit\n")
	fmt.Scanf("%d\n", &selectNum)

	switch {
	case selectNum == 1:
		fmt.Printf("Now add new student: \n")
		AddStuInfo()
	case selectNum == 2:
		fmt.Printf("Now modify student info: \n")
		ModifyUserInfo()
		break
	case selectNum == 3:
		fmt.Print("Currently the student info are:\n")
		fmt.Print("NO\tName\tGrade\tGender\tScore\n")
		for k,v := range(student){
			fmt.Printf("%d\t%v\t%v\t%v\t%v\n",k,v.Username,v.Grade,v.Gender,v.Score)
		}
		break
	case selectNum == 4:
		flag = true
		break
	default:
		break
	}

}

func main(){
	for {
		run()
		if flag {
			break
		}
	}
}