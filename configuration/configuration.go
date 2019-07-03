package configuration

import (
	"io/ioutil"
	"fmt"
	"github.com/ghodss/yaml"
)

func LoadConfiguration(v interface{}, filePath string) (*interface{}, error){
	b, err := ioutil.ReadFile(filePath) // just pass the file name
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
