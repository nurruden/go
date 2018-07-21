package main

import(
	"fmt"
)

func calc(op func(args...int)int,op_args...int) int {
	result := op(op_args...)
	fmt.Printf("result = %d\n",result)
	return result
}

func Add(a ...int) int{
	fmt.Printf("func args count:%d\n",len(a))
	var sum int
	for index,arg := range a{
		fmt.Printf("arg[%d]=%d\n",index,arg)
		sum = sum + arg

	}
	return sum
}

func main(){
	calc(func(args...int)int{
		var sum int
		for i:=0; i<len(args);i++{
			sum = sum + args[i]
		}
		return sum

},1,2,3,4,5)
	calc(func(args...int)int{
		var sum int
		for i:=0; i<len(args);i++{
			sum = sum - args[i]
		}
		return sum

	},1,2,3,4,5)

	calc(Add,123)
}