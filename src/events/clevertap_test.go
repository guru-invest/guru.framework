package events

import (
	"fmt"
	"testing"
)

func TestClevertapEvent(t *testing.T) {
	tracker := NewEventTracker("https://us1.api.clevertap.com", "TEST-8W6-4W8-8R6Z", "SCC-AKW-CWUL")

	properties := make(map[string]interface{})
	properties["name"] = "Weverton"
	properties["email"] = "tom@guru.com.vc"
	properties["nickname"] = "Tom"
	properties["date of birth"] = "07/11/1988"

	err := tracker.SendEvent("testeAgoraVai3", "Event", properties)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
