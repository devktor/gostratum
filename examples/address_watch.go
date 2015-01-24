package main

import (
    "github.com/devktor/gostratum"
    "os"
    "fmt"
    "bufio"
)

func main(){
    if(len(os.Args) < 3){
        fmt.Println("usage: stratum_host address")
        os.Exit(1)
    }
    host:=os.Args[1]
    address:=os.Args[2]

    client, err:=gostratum.Connect(host)

    if err != nil{
        fmt.Println("failed to connect: ",err)
        os.Exit(2)
    }

    fmt.Println("enter to exit")

    client.AddressSubscribe(address, func(address string, status string, err error){
        fmt.Println("address update: ",address, " status=",status, " error=",err)
    })

    bufio.NewReader(os.Stdin).ReadString('\n')

}
