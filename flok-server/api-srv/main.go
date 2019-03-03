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
		Port        string `json:"api"`
		AuthPort    string `json:"auth"`
		UserPort    string `json:"user"`
		ProjectPort string `json:"project"`
		PortalPort  string `json:"portal"`
		ChatPort    string `json:"chat"`
		FilePort    string `json:"file"`
		MeetingPort string `json:"meeting"`
		FormPort    string `json:"form"`
	}{}

	json.Unmarshal(data, d)
	// log.Printf("Port: %+v", d)

	a := newApp(d.Port, "./config.json", "localhost"+d.AuthPort, "localhost"+d.UserPort, "localhost"+d.ProjectPort, "localhost"+d.ChatPort, "localhost"+d.PortalPort, "localhost"+d.MeetingPort, "localhost"+d.FilePort, "localhost"+d.FormPort)

	if err := a.initApp(); err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("Exiting")

	a.closeApp()
}
