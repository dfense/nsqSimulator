package main

import (
	"sync"
	"time"

	"github.com/bitly/go-nsq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	clientName    = kingpin.Flag("clientName", "name of client.").Short('c').Required().String()
	ttlPackets    = kingpin.Flag("ttlPackets", "total number of messages to read").Default("1").Int()
	timeBetween   = kingpin.Flag("timeBetween", "number of millis between message reads").Default("1").Int()
	stayConnected = kingpin.Arg("stayConnected", "stay connected after ttlPacketsRead").Bool()
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()

	log.SetLevel(log.DebugLevel)
	log.Debugf("Client Name: %s\n", *clientName)

	readPackets := 0

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// create a new channel on the "write_test" topic
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("creeSim", *clientName, config)

	// service all messages from channel here. Notification
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("ID: %d %v\n", message.Timestamp, message.ID)
		time.Sleep(time.Second * 2)
		readPackets++
		if readPackets >= *ttlPackets {

			if !*stayConnected {
				wg.Done()
			}
			return nil
		}
		return nil
	}))

	// connect to local NSQ
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()

}
