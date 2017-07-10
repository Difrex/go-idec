package idec

import (
	"encoding/base64"
	"errors"
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

	tags, err := ParseTags(txtMessage[0])

	m.Tags = tags
	m.Echo = txtMessage[1]
	m.Timestamp = ts
	m.From = txtMessage[3]
	m.Address = txtMessage[4]
	m.To = txtMessage[5]
	m.Subg = txtMessage[6]
	m.Body = body

	return m, err
}

// parseTags parse message tags and return Tags struct
func ParseTags(tags string) (Tags, error) {
	var t Tags

	if !strings.Contains(tags, "ii/") {
		e := errors.New("Bad tagstring")
		return t, e
	}

	tagsSlice := strings.Split(tags, "/")
	if len(tagsSlice) < 4 {
		t.II = tagsSlice[1]
		return t, nil
	}
	t.II = tagsSlice[1]
	t.Repto = tagsSlice[3]
	return t, nil
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
