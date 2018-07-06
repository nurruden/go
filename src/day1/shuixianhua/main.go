package main
import (
    "fmt"
    "strconv"
)
func main() {
	var a, b int64
	var flag bool = false
    fmt.Println("请输入求a,b：")
    fmt.Scanf("%d,%d", &a, &b)
    if a > b {
        a, b = b, a
    }
    for i := a; i <= b; i++ {
        if isNarcissusFew(i) {
			fmt.Println(i)
			flag = true
		} 
		}
	if flag == false {
		fmt.Println("no")
	}
    
}
func isNarcissusFew(number int64) bool {
    var sum int64
    numberStr := strconv.FormatInt(number, 10)
    for _, data := range numberStr {
        lenght := len(numberStr)
        num := (int64)(data - 48)
        var sum1 int64 = 1
        for lenght != 0 {
            sum1 *= num
            lenght--
        }
        sum += sum1
    }
    if sum == number {
        return true
    }
    return false
}