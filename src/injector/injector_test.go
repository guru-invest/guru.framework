package injector

import (
	"testing"
)

type TesteStructInterface interface {
	ShowName() string
}

type testeStruct struct {
	Name string
	Test int
}

func (t testeStruct) ShowName() string {
	return t.Name
}

func TestRegistration(t *testing.T) {

	var intf TesteStructInterface
	intf = testeStruct{
		Name: "Valor 1",
		Test: 1,
	}

	GetGlobalContainer().Register("exemplo_Tiago", intf)
	TestGetInstance(t)
}

func TestGetInstance(t *testing.T) {
	instance := GetGlobalContainer().MustGet("exemplo_Tiago").(TesteStructInterface)
	println(instance.ShowName())
}
