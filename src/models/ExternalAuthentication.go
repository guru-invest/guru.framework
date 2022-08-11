package models

var EXTERNALAUTHENTICATION ExternalAuthentication = ExternalAuthentication{}

type ExternalAuthentication struct {
	GenialToken string `json:"genial_token"`
	B3Token     string `json:"b3_token"`
}
