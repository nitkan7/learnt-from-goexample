package main

import(
"fmt"
)


//variadic functions:
//functions that can be called with variable number of arguments

//fmt.Println itself is a variadic function

//only the last parameter of a function can be variadic
func main(){
slice:=[]string{"append","appending"}
fmt.Println("slice before variadic call",slice)
variadic(slice...)
fmt.Println("slice after variadic call",slice)

}

func variadic(str ...string){
//fmt.Println(str)
//str=append(str,"appended")
str[0]="appended"
str=append(str,"append")
}
