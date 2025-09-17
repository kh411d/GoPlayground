package main

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	descriptor := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", "root", "password", "localhost", "okadoc")
	db, err := sqlx.Connect("mysql", descriptor)
	if err != nil {
		log.Fatalf("error connecting to DB: %s", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	wg := &sync.WaitGroup{}
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			sql := `INSERT INTO notification_log_sms (appointment_number,phone,message,notification_sms_vendor_id,provider,sent_time,response) 
VALUES ('OKDC-TEST-44','+6289673410001','',68,'nmc-wisoft-solutions',NOW(),'{\"MessageCount\":2}');`
			db.Exec(sql)
		}(wg)
		wg.Wait()
	}
}
