package main
import "fmt"

type PersonInfo struct{
    ID string
    Name string
    Address string
}

func main(){
    var personDB map[string] PersonInfo
    personDB = make(map[string] PersonInfo)

    personDB["12345"] = PersonInfo{"123456", "tom", "Room203,..."}
    personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
    fmt.Println(personDB)

    person, ok := personDB["1234"]
    if ok {
        fmt.Println("Found person", person.Name, "with ID 1234.")
    }else {
        fmt.Println("Did not find person with ID 1234.")
    }
}
