package main

import "fmt"
import "time"


func main(){

fmt.Println("simple switching")
i:=2
fmt.Print("Write ",i," as")
switch i{
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
default:
	fmt.Println("the number has'nt been coded")

}

fmt.Println("multiple expressions in a single case statement")

switch time.Now().Weekday() {
case time.Saturday,time.Sunday :
	fmt.Println("It's a Weekend")
default	:	
	fmt.Println("It's a weekday")


}

t:=time.Now()
switch{
case t.Hour() < 12 : 
	fmt.Println("It's forenoon ")
default	:	
	fmt.Println("It's afternoon ")
}

}

