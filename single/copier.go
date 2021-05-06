package main

import (
	"github.com/jinzhu/copier"
"fmt"
)
type User struct {
  Name string
  Address string
  Phone int
  Unknown string
}

type Citizen struct {
  ID int
  SocialID int
  Name string
  ParentName string
  Address string
  Phone *int
  HomeAddress string
}

func main(){
   var user = &User{
	   Name: "aowow",
	   Address:"asdfasdfasdfas",
	   Phone: 92830234928,
	   Unknown: "this is unknown data",
   }
   c :=  new(Citizen)

   copier.Copy(&c, &user)
 
   fmt.Printf("%#v  Phone:%d \n",c,*c.Phone)
   
}
