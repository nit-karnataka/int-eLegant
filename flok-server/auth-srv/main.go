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
		Port       string `json:"api"`
		AuthPort   string `json:"auth"`
		DevicePort string `json:"device"`
		UserPort   string `json:"user"`
		HubPort    string `json:"hub"`
	}{}

	json.Unmarshal(data, d)

	a := newApp(d.AuthPort, "test", "localhost:27017", "localhost:6379")
	if err := a.init(); err != nil {
		log.Panicf("init error %+v", err)
	}

	go a.listen()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	a.close()
	/* sess, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Panic(err)
	}
	defer sess.Close()

	fmt.Println(time.Unix(1542902763, 0))

	_, err = sess.DB("").C("test").Upsert(bson.M{"_id": bson.ObjectIdHex("5bf6d3eb7df5e57dfe0407e7").String()}, &proto.Profile{
		Email:   "dev@tets.com",
		Id:      bson.ObjectIdHex("5bf6d3eb7df5e57dfe0407e7").String(),
		Updated: time.Now().Unix(),
		Address: "HELLO",
	})

	if err != nil {
		log.Panic(err)
	} */
	/* b, err := json.Marshal(&proto.Profile{
		Id: bson.NewObjectId().Hex(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) */
}
