package main

import (
        "fmt"
        "math/rand"
        "strconv"
        "strings"
        "time"
)

var randomgen *rand.Rand

func srandom() {
        randomgen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func random() int {
        return randomgen.Int()
}

func generateRefNumber2(kueID int64) string {
        randomByTime := time.Now().Format("20060102150405.9999Z07")
        result := strings.Replace(randomByTime, ".", "", -1)
        result = strings.Replace(result, "+", "", -1)
        result = strings.Replace(result, "Z", "", -1)

        return result + fmt.Sprintf("%06d", kueID)
}

func main() {
        //        var rn int
        fmt.Println(time.Now().Unix())
        fmt.Println(strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int()))

        rands := fmt.Sprintf("%d%d", time.Now().Unix(), rand.New(rand.NewSource(time.Now().UnixNano())).Intn(201))
        fmt.Println("rands")
        fmt.Println(rands)

        //srandom() // initialize random number generator
        fmt.Println(generateRefNumber2(2))

        /*
           for i := 0; i < 20; i++ {
                   for {
                           rn = random() % 100000000
                           if rn > 10000000 { break }
                   }
                   fmt.Printf("%d\n", rn)
           }
        */

}
