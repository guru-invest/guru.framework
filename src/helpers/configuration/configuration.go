package configuration

import (
	"encoding/base64"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	http_connector "github.com/guru-invest/guru.framework/src/infrastructure/http-connector"
)

func Load(url string) []byte {
	m := make(map[string]interface{})
	config := getConfig(url)
	kv := kvPair{}
	if err := json.Unmarshal(config, &kv); err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while in parse config json")
	}

	if err := json.Unmarshal(kv.decodeString(), &m); err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while decoding config")
	}
	ret, _ := json.Marshal(m)
	return ret
}

type kvPair []struct {
	LockIndex   int    `json:"LockIndex"`
	Key         string `json:"Key"`
	Flags       int    `json:"Flags"`
	Value       string `json:"Value"`
	CreateIndex int    `json:"CreateIndex"`
	ModifyIndex int    `json:"ModifyIndex"`
}

func (kv kvPair) decodeString() []byte {
	rawString, err := base64.StdEncoding.DecodeString(kv[0].Value)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while decoding config")
	}
	return rawString
}

func getConfig(url string) []byte {
	client := http_connector.HttpClient{}
	_, res, err := client.Get(url)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while update companies")
	}
	return res
}
