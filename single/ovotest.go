package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"

    "github.com/kh411d/ovo"
)

func dBConn(dsn string) *sql.DB {
    conn, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Println("Conn Error: ", err)
        os.Exit(1)
    }
    conn.SetMaxIdleConns(100)
    conn.SetMaxOpenConns(100)
    return conn
}

func validateAuth(mmsdk *ovo.MatahariMall) {
    ovoReq := &ovo.Request{
        CustomerID: 2015449,
        Phone:      "0818228467",
    }
    errM := mmsdk.ValidateOvoIDAndAuthenticateToOvo(ovoReq)

    if errM != nil {
        fmt.Printf("ERRORRRRR %#v\n\n", errM)
    }

    fmt.Printf("OVOREQUEST %#v\n\n", ovoReq)
}

func checkAuth(mmsdk *ovo.MatahariMall) {

    errM := mmsdk.CheckOvoStatus(2015449)

    if errM != nil {
        fmt.Printf("MESSAGE %#v\n", ovo.GetErrMsg(errM))
        fmt.Printf("CODE %v\n", ovo.GetErrCode(errM))
    }

}

func main() {

    db := dBConn("MMSTGDB:R3FASwMt@tcp(localhost:3307)/MMDBV3?charset=utf8&parseTime=True&loc=Local")
    //apiKey := "4be75d4ba67a1652297a313aec3b9a03509de17c93692dffc6104f682330f571"
    //appID := "matahari-mall"
    baseURL := "https://dududev.ovo.id/loyalty-back"
    apiKey := "084b13ecac81e1a8caf1775ad02bd5fa40e7219c8956dba11429a497a0e4cd89"
    appID := "hypermart"
    merchantID := "2"
    ovoClient := ovo.New(baseURL, apiKey, appID, merchantID)
    mmsdk := ovoClient.GetMMsdk(db)

    // params := map[string]string{
    //     "fullname": "khalid",
    //     "phone":    "0818228467",
    // }
    // y, _ := ovoClient.CreateCustomerLinkage("8000848664340600",)
    // fmt.Println(string(y))

    //validateAuth(mmsdk)
    //checkAuth(mmsdk)

    //data, xx := ovoClient.CreateCustomerLinkage("2015449", params)

    bol, err := mmsdk.IsLinkageVerified(2015449)
    fmt.Println(bol)
    fmt.Println(err)

    /*errc := http.ListenAndServe(":3300", nil) // setting listening port
      if errc != nil {
          log.Fatal("ListenAndServe: ", err)
      }*/
}
