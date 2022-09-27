package main

import "fmt"

func main(){
    number := 9
    i:=1
    for {
        if(i>10){
            break;
        }
        fmt.Println(number," X ",i," = ",number*i)
        i++
    }
}
