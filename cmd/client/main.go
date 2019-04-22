package main

import (
	"sync"
	"time"

	"github.com/bitly/go-nsq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	channel       = kingpin.Flag("channel", "name of channel.").Short('c').Required().String()
	topic         = kingpin.Flag("topic", "name of topic.").Short('t').Required().String()
	ttlPackets    = kingpin.Flag("ttlPackets", "total number of messages to read").Default("1").Int()
	timeBetween   = kingpin.Flag("timeBetween", "number of millis between message reads").Default("1").Int()
	stayConnected = kingpin.Flag("stayConnected", "stay connected after ttlPacketsRead").Default("false").Bool()
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()

	log.SetLevel(log.DebugLevel)
	log.Debugf("Channel Name: %s\n", *channel)
	log.Debugf("StayConnected %t\n", *stayConnected)

	readPackets := 0

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// create a new channel on the "write_test" topic
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer(*topic, *channel, config)

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
