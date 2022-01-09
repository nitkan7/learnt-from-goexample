package main

import "fmt"

func main(){

//single variable with type
var a int 
a=10
fmt.Println(a)

//mutliple variables with type
var b,c int = 20,30
fmt.Println(b,c)

//multiple variables without type
var d,e = 40, 50
fmt.Println(d,e)


//mutilple shorthand assignment
f,g := "apple","donkey"
fmt.Println(f,g)

f,h:="orange","monkey"
fmt.Println(f,h)

var i  int
fmt.Println(i)

}
