package main

import "fmt"

func hofe(division func(a,b int))func(a,b int){
	return division
}

func divi(a,b int){
	fmt.Println(a-b) 
}
func ap()int{
	return 0
}
func main(){
	b:= ap()
	fmt.Println(b)
	fmt.Println("haha")
}