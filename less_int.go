package main
import "fmt"

type Integer int

func (a Integer)Less(b Integer)(bool){
    return a < b
}

func (a *Integer)Add(b Integer){
    *a = *a + b
}

func main(){
    var a Integer = 1
    if a.Less(2){
        fmt.Println(a, "less 2")
    }
    a.Add(2)
    fmt.Println("a=", a)
}
