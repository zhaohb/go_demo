package main

import (
    "fmt"
)

type USB interface{
    Name() string
    Connecter
}

type Connecter interface{
    Connect()
}

type PhoneConnecter struct{
    name string
}

func (pc PhoneConnecter) Name() string{
    return pc.name
}

func (pc PhoneConnecter) Connect() {
    fmt.Println("Connected:", pc.name)
}

func main(){
    var a USB
    a = PhoneConnecter{"PhoneConnecter"}
    a.Connect()
    fmt.Println("name:", a.Name())
    Disconnect(a)
}

func Disconnect(usb interface{}){
    switch v := usb.(type){
        case PhoneConnecter:
            fmt.Println("DisConnected:", v.name)
        default:
            fmt.Println("Unknow device.")
    }
}
