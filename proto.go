package idec

// Base IDEC protocol implementation

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// IDEC Extensions. see: https://ii-net.tk/idec-doc/?p=extensions
const (
	listTXT       = "list.txt"
	blacklistTXT  = "blacklist.txt"
	features      = "x/features"
	xcount        = "x/c/"
	echoSchema    = "u/e/"
	messageSchema = "u/m/"
)

// Extensions IDEC extensions
type Extensions struct {
	ListTXT      string `json:"list_txt"`
	BlacklistTXT string `json:"backlist_txt"`
	Features     string `json:"features"`
	XCount       string `json:"xcount"`
}

// NewExtensions ...
func NewExtensions() Extensions {
	e := Extensions{
		listTXT,
		blacklistTXT,
		features,
		xcount,
	}
	return e
}

// FetchConfig node, echo, and other connection settings
type FetchConfig struct {
	Node   string   `json:"node"`
	Echoes []string `json:"echo"`
	Num    int      `json:"count"`
	Offset int      `json:"offset"`
	Limit  int      `json:"limit"`
}

// ID ...
type ID struct {
	Echo  string `json:"echo"`
	MsgID string `json:"msgids"`
}

// GetMessagesIDS get message ids from node
func (f FetchConfig) GetMessagesIDS() ([]ID, error) {
	var ids []ID

	var getURI string
	getEchoes := strings.Join(f.Echoes, "/")

	// Make strings
	offset := strconv.Itoa(f.Offset)
	limit := strconv.Itoa(f.Limit)

	getURI = strings.Join([]string{f.Node, echoSchema, getEchoes, "/", offset, ":", limit}, "")

	// Get messages ids
	response, err := http.Get(getURI)
	if err != nil {
		return ids, err
	}

	defer response.Body.Close()
	c, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ids, err
	}

	var i ID
	var curEcho string
	rawIDS := strings.Split(string(c), "\n")
	for _, line := range rawIDS {

		// Match echoarea
		if strings.Contains(line, ".") {
			curEcho = line
			continue
		}

		// Match message ID
		if !strings.Contains(line, ".") && !strings.Contains(line, ":") && line != "" {
			i.Echo = curEcho
			i.MsgID = line
			ids = append(ids, i)
		}
	}

	return ids, nil
}

// GetAllMessagesIDS get all message ids from node
func (f FetchConfig) GetAllMessagesIDS() ([]ID, error) {
	var ids []ID

	var getURI string
	getEchoes := strings.Join(f.Echoes, "/")

	getURI = strings.Join([]string{f.Node, echoSchema, getEchoes}, "")

	// Get messages ids
	response, err := http.Get(getURI)
	if err != nil {
		return ids, err
	}

	defer response.Body.Close()
	c, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ids, err
	}

	var i ID
	var curEcho string
	rawIDS := strings.Split(string(c), "\n")
	for _, line := range rawIDS {

		// Match echoarea
		if strings.Contains(line, ".") {
			curEcho = line
			continue
		}

		// Match message ID
		if !strings.Contains(line, ".") && !strings.Contains(line, ":") && line != "" {
			i.Echo = curEcho
			i.MsgID = line
			ids = append(ids, i)
		}
	}

	return ids, nil
}

// MSG ...
type MSG struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

// GetRawMessages get messages from node
func (f FetchConfig) GetRawMessages(ids []ID) ([]MSG, error) {
	var messages []MSG

	var messagesIDS []string
	for _, id := range ids {
		messagesIDS = append(messagesIDS, id.MsgID)
	}
	getMessages := strings.Join(messagesIDS, "/")

	getURI := strings.Join([]string{f.Node, messageSchema, getMessages}, "")

	// Get messages ids
	response, err := http.Get(getURI)
	if err != nil {
		e := errors.New(strings.Join([]string{"Failed to get", getURI, "."}, " "))
		return messages, e
	}

	defer response.Body.Close()
	c, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return messages, err
	}

	for _, m := range strings.Split(string(c), "\n") {
		if len(m) == 0 {
			break
		}
		message := strings.Split(m, ":")
		if len(message) > 1 {
			messages = append(messages, MSG{message[1], message[0]})
		}
	}

	return messages, err
}

// Echo echo description
type Echo struct {
	Name        string `json:"name"`
	Size        int    `json:"size"`
	Description string `json:"description"`
}

// GetEchoList ...
func (f FetchConfig) GetEchoList() ([]Echo, error) {
	var echoes []Echo

	// Check node features support
	fres, err := http.Get(strings.Join([]string{f.Node, features}, "/"))
	if err != nil {
		return echoes, err
	}
	defer fres.Body.Close()

	c, err := ioutil.ReadAll(fres.Body)
	if err != nil {
		return echoes, err
	}

	if !strings.Contains(string(c), listTXT) {
		err = errors.New("Node does not support echoes list")
		return echoes, err
	}

	lres, err := http.Get(strings.Join([]string{f.Node, listTXT}, "/"))
	if err != nil {
		return echoes, err
	}
	defer lres.Body.Close()

	l, err := ioutil.ReadAll(lres.Body)
	if err != nil {
		return echoes, err
	}

	echoes, err = ParseEchoList(string(l))

	return echoes, err
}
