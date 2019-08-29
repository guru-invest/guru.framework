package configuration

import (
	"fmt"
	"testing"
)

func TestReadYaml(t *testing.T) {
	config := LoadYmlFile("test_configuration.yml")
	fmt.Println(config["key"].(string))
}

func TestWriteYaml(t *testing.T) {
	config := LoadYmlFile("test_configuration.yml")
	SaveYmlFile(config, "test_configuration2.yml")
	fmt.Println(config["key"].(string))
}
