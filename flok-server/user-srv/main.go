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
		Port string `json:"user"`
	}{}

	json.Unmarshal(data, d)

	a := newApp(d.Port, "test", "localhost:27017")
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

	err = sess.DB("").C("test").Insert(&proto.Profile{
		Email:   "dev@tets.com",
		Id:      bson.NewObjectId().String(),
		Updated: time.Now().Unix(),
		Address: "dsjfhj",
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
