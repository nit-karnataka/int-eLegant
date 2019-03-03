package main

import (
	"bytes"
	"encoding/json"
	"flok-server/lib"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func configInit() (string, error) {
	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, map[string]interface{}{
		"jwt_secret": lib.RandStringRunes(32),
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

var tpl = template.Must(template.New("initial-config").Parse(strings.TrimSpace(`{
	"jwtSecret": "{{.jwt_secret}}"
}
`)))

func getConfig(configFile string) *config {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {

		log.Printf("creating initial configuration: %s", configFile)

		cfg, err := configInit()
		if err != nil {
			log.Fatalf("failed to generate initial configuration: %s", err)
		}

		if err != nil {
			log.Fatalf("failed to generate initial configuration: %s", err)
		}

		err = ioutil.WriteFile(configFile, []byte(cfg), 0666)
		if err != nil {
			log.Fatalf("failed to write configuration file: %s", err)
		}
	}

	data, e := ioutil.ReadFile(configFile)
	if e != nil {
		log.Fatalf("failed to load config: %v", e)
	}

	cfg := &config{}
	err := json.Unmarshal(data, cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal config: %v", e)
	}
	return cfg
}
