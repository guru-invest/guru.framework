package commons

import shortid "github.com/jasonsoft/go-short-id"

func GenerateShortId(characters int) string{
	opt := shortid.Options{
		Number:        characters,
		StartWithYear: false,
		EndWithHost:   false,
	}
	id := shortid.Generate(opt)
	return id
}
