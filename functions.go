package main

import "fmt"

func main(){
out:=adder(1,2)
fmt.Println("The sum of the 2 numbers is:",out)
}

//two integer params and returns an integer
func adder(a int, b int)int{
return a+b
}

//func adder(a,b int) can also be used
