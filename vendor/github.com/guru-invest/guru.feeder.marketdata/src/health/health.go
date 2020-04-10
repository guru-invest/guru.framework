package health

import "time"

type Health struct {
	ExecutionStatus               string `json:"execution_status,omitempty"`
	AverageExecutionTime          string `json:"average_execution_time,omitempty"`
	RequestAverageExecutionTime   string `json:"request_average_execution_time,omitempty"`
	TopMoversAverageExecutiontime string `json:"movers_average_execution_time,omitempty"`
	StocksCount                   int    `json:"stocks_count,omitempty"`
	MoversCount                   int    `json:"movers_count,omitempty"`
	ConnectionStatus              int    `json:"connection_status,omitempty"`
}

var HEALTH Health

func (h *Health) SetExecutionStatus(elapsed time.Duration) {
	if elapsed <= (100 * time.Millisecond) {
		//amazing
		h.ExecutionStatus = "AMAZING"
	} else if elapsed >= (time.Duration(101)*time.Millisecond) && elapsed <= (time.Duration(1100)*time.Millisecond) {
		//good
		h.ExecutionStatus = "GOOD"
	} else if elapsed >= (time.Duration(1101)*time.Millisecond) && elapsed <= (time.Duration(1800)*time.Millisecond) {
		//warning
		h.ExecutionStatus = "WARNING"
	} else if elapsed >= (time.Duration(1801)*time.Millisecond) && elapsed <= (time.Duration(2500)*time.Millisecond) {
		//slow
		h.ExecutionStatus = "SLOW"
	} else if elapsed >= (time.Duration(2501) * time.Millisecond) {
		//problem in source
		h.ExecutionStatus = "TOO SLOW"
	}
}
