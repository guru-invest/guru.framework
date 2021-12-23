package injector_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/guru-invest/guru.framework/src/injector"
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

	injector.GetGlobalContainer().Register("exemplo_Tiago", intf)
}

func TestDeregister(t *testing.T) {
	TestRegistration(t)
	injector.GetGlobalContainer().Deregister("exemplo_Tiago")

}

func TestHas(t *testing.T) {
	TestRegistration(t)
	instance := injector.GetGlobalContainer().Has("exemplo_Tiago")

	if !instance {
		t.Log("Has method should return true")
		t.Fail()
	}
}

func TestMustGet(t *testing.T) {
	TestRegistration(t)
	instance := injector.GetGlobalContainer().MustGet("exemplo_Tiago").(TesteStructInterface)

	println(instance.ShowName())
}

func TestMustGetWithoutContainer(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestMustGetWithoutContainer should panic")
		}
	}()

	injector.GetGlobalContainer().Deregister("exemplo_Tiago")
	injector.GetGlobalContainer().MustGet("exemplo_Tiago")
}

func TestInvoke(t *testing.T) {
	TestRegistration(t)
	fn := func(arg1 TesteStructInterface, srg2 bool) { fmt.Println("xpto") }
	injector.GetGlobalContainer().Invoke("exemplo_Tiago", fn)

}

func TestInvokeInvalidFunc(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestInvokeInvalidFunc should panic")
		}
	}()

	injector.GetGlobalContainer().Invoke("exemplo_Tiago", "anything but func")

}

func TestMustInvoke(t *testing.T) {
	TestRegistration(t)
	fn := func(arg1 TesteStructInterface) { fmt.Println("xpto") }
	injector.GetGlobalContainer().MustInvoke("exemplo_Tiago", fn)

}

func TestMustInvokeInvalidFunc(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestMustInvokeInvalidFunc should panic")
		}
	}()

	injector.GetGlobalContainer().MustInvoke("exemplo_Tiago", "anything but func")

}

func TestMustInvokeMany(t *testing.T) {
	TestRegistration(t)
	fn := injector.GetGlobalContainer().MustInvokeMany("exemplo_Tiago")

	if reflect.TypeOf(fn).Kind() != reflect.Func {
		t.Log("MustInvokeMany method should return a reflect.Func")
		t.Fail()
	}

	fn(func(arg1 TesteStructInterface) { fmt.Println("xpto") })
}

func TestMustInvokeManyInvalidFunc(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestMustInvokeManyInvalidFunc should panic")
		}
	}()

	TestRegistration(t)
	fn := injector.GetGlobalContainer().MustInvokeMany("exemplo_Tiago")
	fn("anything but func")
}
