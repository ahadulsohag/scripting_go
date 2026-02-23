package main//Higher order function
import "fmt"
func hofe(a,b int, addition func(a,b int)int)int{
	res:= a*b+addition(a,b)
	return res
}

func add (a int,b int)int{
	return a+b
}

func main(){
	res:=hofe(1,2,add)
	fmt.Println(res)
}
