package main

import (
	"flag"
	"log"

	nsq "github.com/nsqio/go-nsq"
)

var (
	topic   = flag.String("t", "", "What topic")
	payload = flag.String("p", "", "string payload")
)

func main() {
	flag.Parse()
	//create producer
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("0.0.0.0:4150", config)

	err := w.Publish(*topic, []byte(*payload))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
}
