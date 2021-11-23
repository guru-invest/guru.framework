package events

import (
	"fmt"
	"testing"
)

func TestClevertapEvent(t *testing.T) {
	tracker := NewEventTracker("https://us1.api.clevertap.com", "848-876-5R6Z", "SAE-KMY-UWUL")

	//Profile

	// err := tracker.SendProfile("testeAgoraVai6", "Weverton", "male", "tom+3@guru.com.vc", "+5511947588847", "Tom", "07/11/1988", "São Paulo", "SP", "Brasil")
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// }

	//err := tracker.CreateProfile("fzVzgo8b", "Tiago Sanches 2", "tiago@guru.com.vc", "Tiago Sanches 2")
	//err := tracker.SendUserProperty("fzVzgo8b", "Years Old", "40")
	//if err != nil {
	//	fmt.Printf("%v", err)
	//}

	//Events
	properties := make(map[string]interface{})
	properties["name"] = "ROLINO"

	err := tracker.SendEvent("fzVzgo8b", "pendency_status", properties)
	if err != nil {
		fmt.Printf("%v", err)
	}

}
