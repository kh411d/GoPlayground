package main

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"sort"
	"sync"
)

func SliceFilter(q string, list interface{}, f func(interface{}) string) interface{} {

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	s := reflect.ValueOf(list)
	if s.Kind() != reflect.Slice {
		return nil
	}

	slices := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		slices[i] = s.Index(i).Interface()
	}

	type result struct {
		result []interface{}
		data   map[int][]interface{}
	}

	r := new(result)

	totalBucketItems := 10

	totalBucket := int(math.Ceil(float64(len(slices)) / float64(totalBucketItems)))

	r.data = make(map[int][]interface{})

	checker := func(idxBucket int, q string, s []interface{}, r *result) {
		for _, v := range s {
			re := regexp.MustCompile("(?i).*" + q + ".*")
			matchThisString := f(v)
			if re.MatchString(matchThisString) {
				mutex.Lock()
				r.data[idxBucket] = append(r.data[idxBucket], matchThisString)
				mutex.Unlock()
			}
		}
		wg.Done()
	}

	for i := 1; i <= totalBucket; i++ {
		skip := (i - 1) * totalBucketItems
		limit := skip + totalBucketItems
		if i == totalBucket {
			limit = len(slices)
		}
		temp := slices[skip:limit]

		wg.Add(1)
		go checker(i, q, temp, r)
	}

	wg.Wait()

	keys := make([]int, 0)
	for k, _ := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		r.result = append(r.result, r.data[k]...)
	}

	return r.result
}

func main() {

	var test = []string{"kambing", "guudng", "kuda", "kambing1", "gunung1", "kuda1", "kambing2", "gunung2", "kuda2", "kambing3", "gunung3", "kuda3",
		"kambing", "guudng", "kuda", "kambing1", "gunung1", "kuda1", "kambing2", "gunung2", "kuda2", "kambing3", "gunung3", "kuda3",
		"kambing", "guudng", "kuda", "kambing1", "gunung1", "kuda1", "kambing2", "gunung2", "kuda2", "kambing3", "gunung3", "kuda3",
		"kambing", "guudng", "kuda", "kambing1", "gunung1", "kuda1", "kambing2", "gunung2", "kuda2", "kambing3", "gunung3", "kuda3"}

	var test2 = []map[string]string{
		map[string]string{
			"name": "kambing",
			"nik":  "1a234",
		},
		map[string]string{
			"name": "kambing1",
			"nik":  "1a234",
		},
		map[string]string{
			"name": "kambing2",
			"nik":  "1b234",
		},
		map[string]string{
			"name": "kuda",
			"nik":  "1b234",
		},
		map[string]string{
			"name": "kuda1",
			"nik":  "1c234",
		},
		map[string]string{
			"name": "kuda2",
			"nik":  "1c234",
		},
	}

	x := SliceFilter("ud", test, func(val interface{}) string {
		return val.(string)
	})

	y := SliceFilter("1b", test2, func(val interface{}) string {
		g := val.(map[string]string)
		return g["name"]
	})

	fmt.Println(x)
	fmt.Println(y)
}
