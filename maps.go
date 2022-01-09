package main

import "fmt"

func main(){

//make(map[key]value)

m:=make(map[string]int)

//m[key]=value
m["s1"]=1
m["s2"]=2

fmt.Println("the map:",m)

//to assign the value of the corresponding key to the variable
value:=m["s1"]
fmt.Println(value)

//to find the number of key value pairs present in the map
fmt.Println(len(m))

//to delete a key value pair
delete(m,"s1")
fmt.Println("The new map is:",m)

//to find if a key is present in the map
_,prs:=m["s1"]
fmt.Println(prs)


//one line declaration and initialisation

n:=map[string]int{"foo":1,"bar":2}
fmt.Println(n)

}
