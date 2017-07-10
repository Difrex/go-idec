package idec

import (
	"encoding/base64"
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
	Tags      string `json:"tags"`
	Repto     string `json:"repto"`
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
