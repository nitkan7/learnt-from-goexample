package main

import "fmt"
import "math/rand"

func main(){
//go has support for multiple return values

//this can be useful  at times like returning VALUE and ERROR from a function

val1,val2:=manyreturns()
fmt.Println("The two random numbers that were returned are:")
fmt.Println(val1,val2)
}

// (int,int) denotes the function returns two integers
func manyreturns()(int, int){
return rand.Intn(100),rand.Intn(2000)
}
