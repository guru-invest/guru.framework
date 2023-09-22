package healthcheck

import (
	"net/http"
	"time"
)

type HealthCheckStatus struct {
	Status    string            `json:"status"`
	Component string            `json:"component"`
	Failures  map[string]string `json:"failures"`
	Success   map[string]string `json:"success"`
	Time      time.Time         `json:"time"`
}

//Status default 200 - OK
func (t HealthCheckStatus) New(component string) HealthCheckStatus {
	return HealthCheckStatus{
		Status:    http.StatusText(http.StatusOK),
		Failures:  make(map[string]string),
		Success:   make(map[string]string),
		Time:      time.Now(),
		Component: component,
	}
}

func (t HealthCheckStatus) Check() (int, HealthCheckStatus) {
	if len(t.Failures) > 0 && len(t.Success) == 0 {
		t.Status = http.StatusText(http.StatusServiceUnavailable)
		return http.StatusServiceUnavailable, t
	}

	if len(t.Failures) > 0 && len(t.Success) > 0 {
		t.Status = "Partially Unavailable"
		return http.StatusServiceUnavailable, t
	}

	return http.StatusOK, t
}
