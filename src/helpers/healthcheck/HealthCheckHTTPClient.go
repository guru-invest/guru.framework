package healthcheck

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const defaultRequestTimeout = 5 * time.Second

type HealthCheckHTTPClient struct {
	Properties []HTTPProperty
	Context    context.Context
}

type HTTPProperty struct {
	Name          string
	URI           string
	Authorization string
	HTTPMethod    string
	Timeout       time.Duration
}

/// Chamadas para o HTTP do HealthCheck
/// Aqui a gente passa sempre uma propriedade por vez, porem na implementacao caso exista mais de uma chamada, criar o []HTTPProperty e chamar o New dentro do foreach
func (t HealthCheckHTTPClient) Check(property HTTPProperty) error {
	if property.Timeout == 0 {
		property.Timeout = defaultRequestTimeout
	}

	if property.HTTPMethod == "" {
		property.HTTPMethod = "GET"
	}

	req, err := http.NewRequest(property.HTTPMethod, property.URI, nil)
	if err != nil {
		return fmt.Errorf("creating the request for the health check failed: %w", err)
	}

	ctx, cancel := context.WithTimeout(t.Context, property.Timeout)
	defer cancel()
	if property.Authorization != "" {
		req.Header.Set("Authorization", "bearer "+property.Authorization)
	}
	// Inform remote service to close the connection after the transaction is complete
	req.Header.Set("Connection", "close")
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("making the request for the health check failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("remote service is not available at the moment")
	}

	return nil
}
