package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/cjlapao/common-go-microsoft-teams/interfaces"
)

type MicrosoftTeamsMessage struct {
	Type        string                            `json:"type"`
	Attachments []MicrosoftTeamsMessageAttachment `json:"attachments"`
}

type MicrosoftTeamsMessageAttachment struct {
	ContentType string                       `json:"contentType"`
	ContentURL  interface{}                  `json:"contentUrl"`
	Content     MicrosoftTeamsMessageContent `json:"content"`
}

type MicrosoftTeamsMessageContent struct {
	Schema  string                     `json:"$schema"`
	Type    string                     `json:"type"`
	Version string                     `json:"version"`
	Body    []interfaces.BodyElement   `json:"body"`
	Actions []interfaces.ActionElement `json:"actions,omitempty"`
}

func NewAdaptiveCardMessage(elements []interfaces.BodyElement) *MicrosoftTeamsMessage {
	msg := MicrosoftTeamsMessage{
		Type: "message",
		Attachments: []MicrosoftTeamsMessageAttachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				ContentURL:  nil,
				Content: MicrosoftTeamsMessageContent{
					Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
					Type:    "AdaptiveCard",
					Version: "1.4",
					Body:    elements,
					Actions: make([]interfaces.ActionElement, 0),
				},
			},
		},
	}

	return &msg
}

func GetDateFormat(t time.Time) string {
	return fmt.Sprintf("{{DATE(%s, SHORT)}}", t.Format(time.RFC3339))
}

func (msg MicrosoftTeamsMessage) GetJson() ([]byte, error) {
	if d, err := json.MarshalIndent(msg, "", "  "); err != nil {
		return nil, err
	} else {
		return d, nil
	}
}

func (msg MicrosoftTeamsMessage) Reader() io.Reader {
	if d, err := msg.GetJson(); err != nil {
		return nil
	} else {
		return bytes.NewReader(d)
	}
}

func (msg MicrosoftTeamsMessage) AddAction(action interfaces.ActionElement) error {
	if len(msg.Attachments) != 1 {
		return errors.New("can only have one attachment")
	}

	msg.Attachments[0].Content.Actions = append(msg.Attachments[0].Content.Actions, action)

	return nil
}

func (msg MicrosoftTeamsMessage) AddElement(element interfaces.BodyElement) error {
	if len(msg.Attachments) != 1 {
		return errors.New("can only have one attachment")
	}

	msg.Attachments[0].Content.Body = append(msg.Attachments[0].Content.Body, element)

	return nil
}
