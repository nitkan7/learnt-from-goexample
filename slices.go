package main

import "fmt"

func main(){
s:=make([]string, 3)

fmt.Println(s)

s[0]="a"
fmt.Println(len(s))
s[1]="b"
fmt.Println(len(s))
s[2]="c"
fmt.Println(len(s))

fmt.Println("set:",s)
fmt.Println(s)
fmt.Println("get:",s[0],s[1],s[2])
fmt.Println(len(s))

fmt.Println("use append function to append to the slice...It also returns a slice")
s=append(s,"D")
fmt.Println(s)
fmt.Println(len(s))

s=append(s,"e","f")
fmt.Println(s)

c:= make([]string,len(s))
copy(c,s)
fmt.Println(c)

}
