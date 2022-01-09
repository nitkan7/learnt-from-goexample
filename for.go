package main

import "fmt"
import "time"

func main(){

fmt.Println("for is the only looping statement present in golang but there are lots of ways in which for loop is there")

i:=1
fmt.Println("while like for loop")
//I guess this is something like a while loop
for i<=3 {
fmt.Println(i)
i++
}

fmt.Println("the classical for loop")
for j:=0;j<=i;j++{
fmt.Println(j)
}


fmt.Println("INfinite loop")
for ;;{
fmt.Println("just print and sleep")
time.Sleep(time.Second)
}






}
