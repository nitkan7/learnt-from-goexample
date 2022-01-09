package main

import "fmt"

func main(){
var a [5]int

fmt.Println("a:",a)

a[4]=77
fmt.Println("set:",a)
fmt.Println("get:",a[4],a[0])
fmt.Println("length:",len(a))

b:= [5]int{1,2,3,4,5}
fmt.Println("b has been declared",b)

fmt.Println("two dimensional array example")
var twoD [2][3]int
for i:=0;i<2;i++{
  for j:=0;j<3;j++{
    twoD[i][j]=i+j
  }
}

fmt.Println(twoD)


}
