package main

import(
    "flag"
    "fmt"
    "math/rand"
    "time"
)

// func num(length int) {
//     password := make([]int,length)
//     for i:=0;i<len(password); i++ {
//         rand.Seed(time.Now().UnixNano())
//         number := rand.Intn(9)
//         password[i] = number
  
//     }
// }
// const(
//     numbers = "0123456789"
//     letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//     mixletters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//     specialLetters = "~!@#$%^&*()_+{}|[]<>?"
// )
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
	passwdType := flag.String("t","num","Password type, the default is num,char is english charater, mix is combaine number and english,advance is number,english and specil character ")
    passwdLength := flag.Int("l",6,"Password length, default and at least is 6")
    flag.Parse()
    fmt.Printf("type=%s\n", *passwdType)
    fmt.Printf("length=%d\n",*passwdLength)
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
    default:
        fmt.Println("not a vowel")

    }
    
    
}