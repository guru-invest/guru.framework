package configuration

import (
	b "encoding/base64"
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

func (kv kvPair) decodeString() []byte {
	rawString, err := b.StdEncoding.DecodeString(kv[0].Value)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while decoding config")
	}
	return rawString
}

func getConfig(url string) []byte {
	client := http_connector.HttpClient{}
	res, err := client.Get(url, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while update companies")
	}
	return res
}
