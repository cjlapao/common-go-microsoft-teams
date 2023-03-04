package microsoft_teams_service

import (
	"errors"
	"net/http"
	"strconv"

	log "github.com/cjlapao/common-go-logger"
	"github.com/cjlapao/common-go-microsoft-teams/message"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/guard"
)

const (
	WEBHOOK_URL_ENVIRONMENT_VAR = "microsoft_teams_service__webhook_url"
)

type MicrosoftTeamsService struct {
	context *execution_context.Context
	logger  *log.LoggerService
	Webhook string
}

func New() *MicrosoftTeamsService {
	svc := MicrosoftTeamsService{
		logger:  log.Get(),
		context: execution_context.Get(),
	}

	return &svc
}

func (svc *MicrosoftTeamsService) Send(msg *message.MicrosoftTeamsMessage) error {
	if svc.Webhook == "" {
		svc.Webhook = svc.context.Configuration.GetString(WEBHOOK_URL_ENVIRONMENT_VAR)
	}

	if err := guard.EmptyOrNil(svc.Webhook, "WebHook"); err != nil {
		return errors.New("webhook cannot be empty")
	}

	reader := msg.Reader()
	if err := guard.EmptyOrNil(reader); err != nil {
		svc.logger.Exception(err, "There was an error generating the message body")
		return err
	}

	if r, err := http.NewRequest("POST", svc.Webhook, msg.Reader()); err != nil {
		svc.logger.Exception(err, "There was an error generating the webhook request")
		return err
	} else {
		client := http.DefaultClient
		if r, err := client.Do(r); err != nil {
			svc.logger.Exception(err, "There was an error sending the message to the teams webhook, status code %s", r.Status)
			return err
		} else {
			svc.logger.Info("Message was sent successfully to the teams webhook with status %s", strconv.Itoa(r.StatusCode))
			return nil
		}
	}
}
