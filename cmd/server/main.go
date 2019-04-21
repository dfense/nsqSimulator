package main

import (
	"time"

	"github.com/bitly/go-nsq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	msBetween = kingpin.Flag("ms", "millisecond to wait between sending messages.").Short('m').Required().Int()
	count     = kingpin.Flag("count", "total number of messages to send").Default("1").Int()
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()

	log.SetLevel(log.DebugLevel)
	duration := time.Duration(time.Duration(*msBetween) * time.Millisecond)

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	// loop for 1000 times
	for i := 0; i < *count; i++ {

		err := w.Publish("creeSim", []byte("Cree Test Message"))
		if err != nil {
			log.Panic("Could not connect")
		}

		log.Debugf("Sleeping : %d\n", *count)
		time.Sleep(duration)
	}

	w.Stop()
}
