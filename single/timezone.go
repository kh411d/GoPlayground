package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//test doang lagi
	// then := now.Add(-179 * time.Second)
	// fmt.Println(now)
	// fmt.Println(then)

	now = now.Add(2 * time.Hour)

	//layout := "2006-01-02 15:04:05"
	//2022-06-14 19:20:00

	now, _ = time.Parse("2006-01-02 15:04:05", "2025-05-21 00:00:00")

	loc, _ := time.LoadLocation("Asia/Jakarta")

	t := now.In(loc).Format(time.RFC3339)
	fmt.Printf("Asia/Jakarta:\n%v\n", t)
	fmt.Printf("UTC:\n%v\n", now.UTC().Format(time.RFC3339))

	fmt.Println(now.In(loc).Before(now.UTC()))

	fmt.Println("----test fro offset---")
	localOffset := "-420"
	jobtime := "2022-11-27T13:00:00.000Z"
	//"job_time_utc": "2022-11-27T06:00:00.000Z",
	//"job_date_utc": "2022-11-27T00:00:00.000Z",

	d, _ := time.ParseDuration(localOffset + "m")

	x := GetLocalTime(jobtime, d)
	fmt.Println(x)

	thetime, _ := time.Parse(time.RFC3339Nano, "2022-11-27T06:00:00.000+08:00")
	tzs, _ := thetime.Zone()
	fmt.Printf("loc %s\n", tzs)
}

func GetLocalTime(jobtime string, localOffset time.Duration) time.Time {
	utcOffset := -localOffset

	d := utcOffset

	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	tz := "Z"
	if utcOffset > 0 {
		tz = fmt.Sprintf("+%02d:%02d", h, m)
	} else if utcOffset < 0 {
		tz = fmt.Sprintf("-%02d:%02d", -(h), m)
	}

	t, _ := time.Parse(time.RFC3339Nano, jobtime)

	//RFC3339 formatted "2006-01-02T15:04:05Z07:00"
	DatetimeWithLocalTZ := fmt.Sprintf(
		"%04d-%02d-%02dT%02d:%02d:%02d%s",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), tz,
	)

	dt, _ := time.Parse(time.RFC3339, DatetimeWithLocalTZ)

	return dt
}
