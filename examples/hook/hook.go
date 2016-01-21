package main

import (
	"github.com/sp-marcos-quesada/logrus"
	"gopkg.in/gemnasium/logrus-airbrake-hook.v2"
	"sync"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.TextFormatter{Mutex: &sync.Mutex{}}
	log.Hooks.Add(airbrake.NewHook(123, "xyz", "development"))
}

func main() {
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
