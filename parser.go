package idec

import (
	"encoding/base64"
	"strconv"
	"strings"
)

// ParseMessage ...
func ParseMessage(message string) (Message, error) {
	var m Message
	plainMessage, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return m, err
	}

	txtMessage := strings.Split(string(plainMessage), "\n")

	var body string
	for i := 8; i < len(txtMessage); i++ {
		body = strings.Join([]string{body, txtMessage[i]}, "\n")
	}

	ts, err := strconv.Atoi(txtMessage[2])
	if err != nil {
		return m, err
	}
	m.Tags = txtMessage[0]
	m.Echo = txtMessage[1]
	m.Timestamp = ts
	m.From = txtMessage[3]
	m.Address = txtMessage[4]
	m.To = txtMessage[5]
	m.Subg = txtMessage[6]
	m.Body = body

	println(m.Tags)

	return m, err
}

// ParseEchoList parse /list.txt
func ParseEchoList(list string) ([]Echo, error) {
	var echoes []Echo
	for _, e := range strings.Split(list, "\n") {
		desc := strings.Split(e, ":")
		if len(desc) <= 1 {
			break
		}
		count, err := strconv.Atoi(desc[1])
		if err != nil {
			return echoes, err
		}
		echoes = append(echoes, Echo{desc[0], count, desc[2]})
	}

	return echoes, nil
}
