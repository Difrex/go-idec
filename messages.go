package idec

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"
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

// PointMessage
type PointMessage struct {
	Echo      string `json:"echo"`
	To        string `json:"to"`
	Subg      string `json:"subg"`
	EmptyLine string `json:"empty_line"`
	Repto     string `json:"repto"`
	Body      string `json:"body"`
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

// Make bundle from point message
func (p PointMessage) MakeBundleMessage(from, address string) (*Message, string, error) {
	var result string
	var message *Message

	nodeTime := strconv.Itoa(int(time.Now().Unix()))
	tags := Tags{"ok", p.Repto}

	message.Tags = tags
	message.Echo = p.Echo
	message.Timestamp = int(time.Now().Unix())
	message.From = from
	message.Address = address
	message.To = p.To
	message.Subg = p.Subg
	message.Body = p.Body

	strTags, err := tags.CollectTags()
	if err != nil {
		return message, "", err
	}
	rawMessage := strings.Join([]string{strTags, p.Echo,
		nodeTime, from, address, p.To, p.Subg, p.EmptyLine, p.Body}, "\n")

	result = base64.StdEncoding.EncodeToString([]byte(rawMessage))

	return message, result, nil
}

// PrepareMessageForSend Make base64 encoded message. Client.
func (p PointMessage) PrepareMessageForSend() string {
	var result string

	var rawMessage string
	if p.Repto != "" {
		rawMessage = strings.Join([]string{p.Echo, p.To, p.Subg, p.EmptyLine, p.Repto, p.Body}, "\n")
	}
	rawMessage = strings.Join([]string{p.Echo, p.To, p.Subg, p.EmptyLine, p.Body}, "\n")

	result = base64.StdEncoding.EncodeToString([]byte(rawMessage))

	return result
}
