package stuinfo

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
	// "String"
	// "unsafe"
)

var flagstudent bool = false
type Book struct {
	BookID int `json:"bookid"`
	BookName string `json:"bookname"`
}
// var BookArray[3]Book
    
type Student struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Grade int `json:"grade"`
	Gender string `json:"gender"`
	BookArray []Book `json:"bookarray"`
}

var student []Student = make([]Student,0,0)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func (pstudent *Student)AddStuInfo() {
    
    
	// var stu Student
	fmt.Println("Please input student ID:")
	fmt.Scanf("%d\n",&pstudent.ID)
	fmt.Println("Please input student name:")
	fmt.Scanf("%s\n",&pstudent.Username)
	fmt.Println("Please input student grade:")
	fmt.Scanf("%d\n",&pstudent.Grade)
	fmt.Println("Please input student gender:")
	fmt.Scanf("%s\n",&pstudent.Gender)
    fmt.Print("ID\tName\tGrade\tGender\n")
	fmt.Printf("%v\t%v\t%v\t%v\n",pstudent.ID,pstudent.Username,pstudent.Grade,pstudent.Gender)
	// fmt.Printf("%v\n",*pstudent)
	student = append(student,*pstudent)
	data, err := json.Marshal(student)
	if err != nil{
		fmt.Printf("%v\n",err)
	}
	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        log.Fatal(err)
    }
    defer fp.Close()
    _, err = fp.Write(data)
    if err != nil {
        log.Fatal(err)
    }
       
} 

func ModifyUserInfo() {
    var studentExist bool = false
	fmt.Printf("Please input the name which you want to modify:")
  
	var studentName string
	fmt.Scanf("%v\n", &studentName)
	for k, stu := range student {
		if stu.Username == studentName {

			var name string
			fmt.Printf("Please input the new name:\n")
			fmt.Scanf("%s\n", &name)
			stu.Username = name
			

			fmt.Printf("Please input the new grade\n")
			var grade int
			fmt.Scanf("%v\n", &grade)
			stu.Grade = grade

			fmt.Printf("Please input the new gender\n")
			var gender string
			fmt.Scanf("%v\n", &gender)
			stu.Gender = gender
			student[k] = stu
			// fmt.Printf("%v",stu)
			fmt.Print("ID\tName\tGrade\tGender\n")
	        fmt.Printf("%v\t%v\t%v\t%v\n",student[k].ID,student[k].Username,student[k].Grade,student[k].Gender)
			studentExist = true
			break
		}

	}
	if studentExist == false {
        fmt.Printf("The name is not exist!\n")
	}
	

}
func readStudentinfo() {
	fp, err := os.OpenFile("./data.json", os.O_RDONLY, 0755)
    defer fp.Close()
    if err != nil {
        log.Fatal(err)
    }
    data := make([]byte, 1000)
    n, err := fp.Read(data)
    if err != nil {
        log.Fatal(err)
	}
	var studentslice []Student
	if err := json.Unmarshal(data[:n],&studentslice);err != nil{
		fmt.Printf("%v\n",err)
	}
    for _,v := range(studentslice){
		fmt.Printf("%v\t%v\t%v\t%v\n",v.ID,v.Username,v.Grade,v.Gender)
	}
}
func initStudent(){

	var selectStr string

	fmt.Print("\t\t\t\tWelcome to the student manage system\t\t\t\n")
	fmt.Print("**********************************************************************************************************\n")
	fmt.Print("Please select the NO:\n")
	fmt.Print("a)Add new student information\nb)Modify student information\nc)Display all student information\nd)Quit\n")
	fmt.Scanf("%s\n", &selectStr)

	switch selectStr{
	case "a":
		fmt.Printf("Now add new student: \n")
		var pstudent *Student = new(Student)
		pstudent.AddStuInfo()
	
	case "b":
		fmt.Printf("Now modify student info: \n")
		ModifyUserInfo()
		break
	case "c":
		fmt.Print("Currently the student info are:\n")
		fmt.Print("NO\tName\tGrade\tGender\n")
		// for _,v := range(student){
		// 	fmt.Printf("%d\t%v\t%v\t%v\n",v.ID,v.Username,v.Grade,v.Gender)
		// }
		readStudentinfo()
        
		break
	case "d":
		flagstudent = true
		break
	default:
		break
	}

}
func RunStudent(){
	for {
		initStudent()
		if flagstudent {
			break
		}
	}
}