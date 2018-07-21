package main

import(
    "flag"
    "fmt"
    "math/rand"
    "time"
)

var numbers = []rune("0123456789")
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var mixletters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var advanceletter = []rune("~!@#$%^&*()_+{}|[]<>?0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func num(length int) string {
    b := make([]rune, length)
    for i := range b {
        rand.Seed(time.Now().UnixNano())
        b[i] = numbers[rand.Intn(len(numbers))]
    }
    return string(b)
}   

func advance(length int) string {
    b := make([]rune, length)
    // advanceletter = fmt.Sprintf("%s%s", mixletters, specialLetters)
    for i := range b {
        rand.Seed(time.Now().UnixNano())
        b[i] = advanceletter[rand.Intn(len(advanceletter))]
    }
    return string(b)
} 


func char(length int) string {
    b := make([]rune, length)
    for i := range b {
        rand.Seed(time.Now().UnixNano())
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
    

func mix(length int) string {
    b := make([]rune, length)
    for i := range b {
        rand.Seed(time.Now().UnixNano())
        b[i] = mixletters[rand.Intn(len(mixletters))]
    }
    return string(b)
}

func main() {
	passwdType := flag.String("t","num","Password type, num is just numbers,char is english charater,\n 	mix is combaine number and english character,\n	advance is combine number,english character and specil character ")
    passwdLength := flag.Int("l",6,"Password length, default and at least is 6")
    flag.Parse()
    // passwordLength range should be [6,12]
	if *passwdLength >= 6 && *passwdLength <=12 {
		switch *passwdType {
		case "num":
			numpassword :=num(*passwdLength)
			fmt.Printf("%v\n",numpassword)
		case "char":
			ss := char(*passwdLength)
			fmt.Printf("%v\n",ss)
		case "mix":
			mixpassword := mix(*passwdLength)
			fmt.Printf("%v\n",mixpassword)
		case "advance":
			advpassword := advance(*passwdLength)
			fmt.Printf("%v\n",advpassword)
		// if passwdType not in num,char,mix,advance, print message,exit
		default:
			fmt.Println("Only num, char,mix,advance could be selected, please re-design password type")
	
		}
		// if passwordLength not in the range, print message,exit
	} else {
		fmt.Println("Password length should be between 6 and 12, cannot generate password for you,please try again")
	}
    
    
}