package main

import "fmt"

type User struct {
    name string
}

type PromoType struct {
    typeId string
}

func gogo(params ...interface{}) {
    fmt.Println(len(params))
    //var user User
    /* user, ok := params[0].(User)
       if !ok {
           fmt.Println(ok)
       }
       fmt.Println(user)
       if len(params) > 1 {
           promo_type, ok := params[1].(PromoType)
           if !ok {
               fmt.Println("asdf")
           }
           fmt.Println(promo_type)
       }*/

    //fmt.Println(len(isi))
    // fmt.Println(isi[0])
}

func main() {

    // x := User{name: "asdf"}
    // y := PromoType{typeId: "1"}
    gogo()
}
