package main

import "fmt"

func main(){
//using range over a  slice
slice:=[]int{1,2,3}
// i is index , v is the value itself
for i,val:= range slice{
fmt.Println(i,val)
}


//blank identifier is used because we are not using the index value
sum:=0
for _,val:= range slice{
sum+=val
}
fmt.Println("The sum of the elements in the slice:",sum)

//using range over a map

mapp:=map[string]int{"apple":1,"orange":2,"grape":3}

for key,value:= range mapp{
fmt.Println("The key is",key,"and the value is",value)
}

//using range over unicode code points
for i,c:=range "go"{
fmt.Println("The index is ",i,"the value is",c)
}

}
