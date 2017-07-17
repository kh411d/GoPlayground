package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

func GetLogFile(dir string) (*os.File, error) {
    fmode := os.FileMode(0775)
    basedir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    dir = strings.Trim(dir, " ")
    dir = strings.Trim(dir, "/")

    logDir := fmt.Sprintf("%s/logs/%s", basedir, dir)
    if _, err := os.Stat(logDir); os.IsNotExist(err) {
        if err := os.MkdirAll(logDir, fmode); err != nil {
            log.Panic(err)
            return nil, err
        }
    }

    logFile := fmt.Sprintf("%s/%s.log", logDir, time.Now().Format("20060102"))

    f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, fmode)
    if err != nil {
        log.Panic("error opening file:", err)
        return nil, err
    }

    return f, nil
}

func main() {
    f, err := GetLogFile("kambinggunung/congea/ ")
    fmt.Println(err)
    fmt.Printf("%v", f)
}
