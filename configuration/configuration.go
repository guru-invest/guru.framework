package configuration

import (
	"fmt"
	"github.com/ghodss/yaml"
	yml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func LoadConfiguration(v interface{}, filePath string) (*interface{}, error){
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	err = yaml.Unmarshal(b, &v)
	if err != nil{
		fmt.Print(err)
		return nil, err
	}
	return &v, nil
}

func LoadYmlFile(fileName string) map[string]interface{} {
	m := make(map[string]interface{})
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	err = yml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return m
}

func SaveYmlFile(data interface{}, filename string){
	d, err := yml.Marshal(&data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile( filename, d, 0777)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func LoadYmlFileToJson(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	b, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return b
}


func SaveJsonToYmlFile(data []byte, filename string){
	d, err := yaml.JSONToYAML(data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile( filename, d, 0777)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
