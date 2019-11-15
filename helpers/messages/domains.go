package messages


type description = string

type listDescription struct {
	CustomerInvalid   			description
	CustomerCodeNotFound		description
	CustomerNotFound			description
	ReferralsNotFound 			description
	MissingKey					description
	InvalidFormat				description
}

var DomainMessages = &listDescription{
	CustomerInvalid: "invalid customer.",
	CustomerCodeNotFound: "customer code not found.",
	CustomerNotFound: "customer not found.",
	ReferralsNotFound: "referral not found.",
	MissingKey: "missing key: ",
	InvalidFormat: "invalid format.",
}

func MissingKey(key string) string{
	return DomainMessages.MissingKey + key
}