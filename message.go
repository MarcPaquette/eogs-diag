package main

import (
	"encoding/json"
	"strings"
)

// Message stored diagnostic report information and errors
type Message struct {
	DiagReport string
	Errors     map[string]string
}

// Validate validates the message body
func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	if json.Valid([]byte(msg.DiagReport)) == false {
		msg.Errors["DiagReport"] = "Please ensure your diagnostic report is valid JSON"
	}
	if strings.TrimSpace(msg.DiagReport) == "" {
		msg.Errors["DiagReport"] = "Please paste a Ops Manager Diagnostics Report"
	}

	return len(msg.Errors) == 0
}

// Deliver delivers things
func (msg *Message) Deliver() error {
	return nil
}
