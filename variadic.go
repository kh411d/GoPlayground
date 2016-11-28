package main 
import (
	"fmt"
)

func gogo (tempe string,isi ...bool) {
fmt.Println(tempe)
     fmt.Println(len(isi))
    // fmt.Println(isi[0])
}

func main() {
	gogo("kambing",true)
}