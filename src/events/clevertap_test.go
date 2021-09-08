package events

import (
	"fmt"
	"testing"
)

func TestClevertapEvent(t *testing.T) {
	tracker := NewEventTracker("https://us1.api.clevertap.com", "TEST-8W6-4W8-8R6Z", "SCC-AKW-CWUL")

	//Profile

	err := tracker.SendProfile("testeAgoraVai6", "Weverton", "male", "tom+3@guru.com.vc", "+5511947588847", "Tom", "07/11/1988", "São Paulo", "SP", "Brasil")
	if err != nil {
		fmt.Printf("%v", err)
	}

	// //Events
	// properties := make(map[string]interface{})
	// properties["Name"] = true

	// err = tracker.SendEvent("testeAgoraVai5", "Teste_de_evento", properties)
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// }

}
