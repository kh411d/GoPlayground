package main

import (
	"fmt"
	"sync"
	"time"
)

type docs struct {
	GroupID int64
}

func main() {
	patientDocuments := []docs{
		{
			GroupID: 1,
		},
		{
			GroupID: 2,
		},
		{
			GroupID: 3,
		},
		{
			GroupID: 2,
		},
		{
			GroupID: 4,
		},
	}

	data := map[int64]string{
		1: "satu",
		2: "dua",
		3: "tiga",
	}

	//Accumulate group id from patient document
	documentGroupIds := make(map[int64]string)
	var wgDocGroupID sync.WaitGroup
	for _, d := range patientDocuments {
		_, ok := documentGroupIds[d.GroupID]
		if d.GroupID > 0 && !ok {
			documentGroupIds[d.GroupID] = ""
		}

	}

	for k := range documentGroupIds {
		wgDocGroupID.Add(1)
		go func(wg *sync.WaitGroup, gids map[int64]string, id int64) {
			gids[id] = data[id]
			defer wg.Done()
		}(&wgDocGroupID, documentGroupIds, k)
	}
	wgDocGroupID.Wait()
	time.Sleep(1 * time.Second)
	fmt.Printf("%#v \n", documentGroupIds)
}
