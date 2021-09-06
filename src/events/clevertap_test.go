package events

import (
	"fmt"
	"testing"
)

func TestClevertapEvent(t *testing.T) {
	tracker := NewEventTracker("https://us1.api.clevertap.com", "TEST-8W6-4W8-8R6Z", "SCC-AKW-CWUL")
	properties := make(map[string]interface{})
	properties["nickname"] = "Tom"
	properties["name"] = "Weverton"
	properties["date of birth"] = "07/11/1988"
	err := tracker.SendEvent("teste123", "Teste de evento via Guru Framework", properties)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
