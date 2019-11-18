package messages

import (
	"bytes"
	"text/template"
)

type Message struct{
	text string
	param string
}

func messageInstance(message string, param string) Message {
	return Message{
		text: message,
		param: param,
	}
}

func ParseTemplate(message string, param string) string{
	msg := messageInstance(message, param)
	t, _ := template.New("message").Parse(msg.text)
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, msg)
	return buf.String()
}
