package main

import (
	"flag"
	"fmt"
	"log"

	nsq "github.com/nsqio/go-nsq"
)

var (
	topic   = flag.String("t", "", "What topic")
	channel = flag.String("c", "", "What channel")
)

func main() {
	flag.Parse()

	//create consumer
	//wg := &sync.WaitGroup{}
	//wg.Add(1)

	config := nsq.NewConfig()

	q, _ := nsq.NewConsumer(*topic, *channel, config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("%#v", string(message.Body))
		//wg.Done()
		return nil
	}))
	//err := q.ConnectToNSQD("0.0.0.0:4150")
	err := q.ConnectToNSQLookupd("0.0.0.0:4161")
	if err != nil {
		log.Panic("Could not connect")
	}
	//wg.Wait()

	fmt.Scanln()
}
