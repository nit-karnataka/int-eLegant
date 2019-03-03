package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

func main() {
	data, _ := ioutil.ReadFile("../cloud-config.json")

	d := &struct {
		Port string `json:"file"`
	}{}

	json.Unmarshal(data, d)

	a := newApp(d.Port)
	if err := a.init(); err != nil {
		log.Panicf("init error %+v", err)
	}

	go a.listen()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	a.close()
}
