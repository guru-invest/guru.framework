package dynamic

import shortid "github.com/jasonsoft/go-short-id"

func GenerateShortId() string{
	opt := shortid.Options{
		Number:        6,
		StartWithYear: false,
		EndWithHost:   false,
	}
	id := shortid.Generate(opt)
	return id
}

func GenerateCustomerCode() string{
	opt := shortid.Options{
		Number:        8,
		StartWithYear: false,
		EndWithHost:   false,
	}
	id := shortid.Generate(opt)
	return id
}
