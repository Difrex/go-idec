package idec

import (
	"encoding/base64"
	"errors"
	"strings"
)

// Message IDEC message structure
type Message struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Address   string `json:"address"`
	Echo      string `json:"echo"`
	Subg      string `json:"subg"`
	ID        string `json:"id"`
	Timestamp int    `json:"timestamp"`
	Body      string `json:"body"`
	Tags      Tags   `json:"tags"`
	Repto     string `json:"repto"`
}

// Tags IDEC message tags
type Tags struct {
	II    string `json:"ii"`
	Repto string `json:"repto"`
}

// CollectTags make ii/ok message from Tags
func (t Tags) CollectTags() (string, error) {
	var tagstring string
	if t.II == "" {
		e := errors.New("Wrong ii/ok tag")
		return "", e
	}
	if t.Repto != "" {
		tagstring = strings.Join([]string{"ii", t.II}, "/")
	} else {
		tagstring = strings.Join([]string{"ii", t.II, "repto", t.Repto}, "/")
	}
	return tagstring, nil
}

// PrepareMessageForSend Make base64 encoded message
func PrepareMessageForSend(m Message) string {
	var result string

	var rawMessage string
	if m.Repto != "" {
		rawMessage = strings.Join([]string{m.Echo, m.To, m.Subg, "", m.Repto, m.Body}, "\n")
	}
	rawMessage = strings.Join([]string{m.Echo, m.To, m.Subg, "", m.Body}, "\n")

	result = base64.StdEncoding.EncodeToString([]byte(rawMessage))

	return result
}
