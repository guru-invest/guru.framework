package messages

type keyDescription = string
type keysDescriptions struct {
	CustomerCode keyDescription
	ReferralCode keyDescription
}

//KeysDescriptions : retorna lista de chaves para erros genericos
var KeysDescriptions = &keysDescriptions{
	CustomerCode: "customer_code",
	ReferralCode: "referral_code",
}

type description = string
type listDescription struct {
	CustomerInvalid      description
	CustomerCodeNotFound description
	CustomerNotFound     description
	ReferralsNotFound    description
	MissingKey           description
	InvalidFormat        description
}

var DomainMessages = &listDescription{
	CustomerInvalid:      "invalid customer.",
	CustomerCodeNotFound: "customer code not found.",
	CustomerNotFound:     "customer not found.",
	ReferralsNotFound:    "referral not found.",
	MissingKey:           "missing key: ",
	InvalidFormat:        "invalid format.",
}

func MissingKey(key keyDescription) string {
	return DomainMessages.MissingKey + key
}
