package events

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	clevertap "github.com/kitabisa/go_sdk_clevertap"
)

type EventTracker struct {
	builder *clevertap.Builder
	service *clevertap.Service
	//clevertapBuild clevertap.BuildClevertap
}

func NewEventTracker(cleverTapUrl string, accountId string, passcode string) EventTracker {
	clevertapBuilder := &clevertap.Builder{}
	service := &clevertap.Service{}
	et := EventTracker{
		builder: clevertapBuilder,
		service: service,
	}
	baseUrl, _ := url.Parse(cleverTapUrl)
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	et.builder.SetBuilder(et.service)
	et.builder.SetHTTPClient(httpClient)
	et.builder.SetBaseURL(baseUrl)
	et.builder.SetAccountID(accountId)
	et.builder.SetPasscode(passcode)
	//et.clevertapBuild = et.builder.Build()
	return et
}

func (et *EventTracker) SendEvent(customer_code string, eventName string, eventProperties map[string]interface{}) error {
	builder := et.builder.Build()
	eventProperties["user_id_type"] = "identity"
	cleverTapResponse := &clevertap.Response{}
	err := builder.SendEvent(customer_code, eventName, eventProperties, cleverTapResponse)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendProfile(customer_code string, name string, gender string, email string, phone string, nickname string, date_of_birth string, city string, state string, country string) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Name"] = name
	userProperties["Gender"] = gender
	userProperties["Email"] = email
	userProperties["Nickname"] = nickname
	userProperties["Date of birth"] = date_of_birth
	userProperties["Phone"] = phone
	userProperties["Location"] = city + ", " + state + " - " + country

	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) CreateProfile(customer_code string, name string, email string, nickname string) error {
	err := et.builder.Build().SendProfile(
		customer_code,
		map[string]interface{}{
			"Name":     name,
			"Email":    email,
			"Nickname": nickname,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendName(customer_code string, name string) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Name"] = name
	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendGender(customer_code string, gender string) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Gender"] = gender
	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendEmail(customer_code string, email string) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Email"] = strings.ToLower(email)
	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendNickname(customer_code string, nickname string) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Nickname"] = nickname
	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendDateOfBirth(customer_code string, date_of_birth string) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Date of birth"] = date_of_birth
	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendPhone(customer_code string, phone string) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Phone"] = phone
	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendLocation(customer_code string, city string, state string, country string) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Location"] = city + ", " + state + " - " + country
	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendDateOfRegistry(customer_code string, dateregistry time.Time) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Date of Registry"] = dateregistry
	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}

func (et *EventTracker) SendDateOfPreRegistry(customer_code string, datePreregistry time.Time) error {
	builder := et.builder.Build()

	userProperties := make(map[string]interface{})
	userProperties["Date of Preregistry"] = datePreregistry
	err := builder.SendProfile(customer_code, userProperties)
	if err != nil {
		return err
	}
	return nil
}
