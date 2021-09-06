package events

import (
	"net/http"
	"net/url"
	"time"

	clevertap "github.com/kitabisa/go_sdk_clevertap"
)

type EventTracker struct {
	builder        *clevertap.Builder
	service        *clevertap.Service
	clevertapBuild clevertap.BuildClevertap
}

func NewEventTracker(cleverTapUrl string, accountId string, passcode string) *EventTracker {
	baseUrl, _ := url.Parse(cleverTapUrl)
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	et := EventTracker{
		builder: &clevertap.Builder{},
		service: &clevertap.Service{},
	}
	et.builder.SetBuilder(et.service)
	et.builder.SetHTTPClient(httpClient)
	et.builder.SetBaseURL(baseUrl)
	et.builder.SetAccountID(accountId)
	et.builder.SetPasscode(passcode)
	et.clevertapBuild = et.builder.Build()
	return &et
}

func (et *EventTracker) SendEvent(customer_code string, eventName string, eventProperties map[string]interface{}) error {
	cleverTapResponse := &clevertap.Response{}
	err := et.clevertapBuild.SendEvent(customer_code, eventName, eventProperties, cleverTapResponse)
	if err != nil {
		return err
	}
	return nil
}
