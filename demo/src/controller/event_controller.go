package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/events"
)

type EventController struct {
}

func (s *EventController) Handle(ctx contracts.Context) (interface{}, error) {
	params := make(map[string]interface{})
	payload := &contracts.Payload{
		Route:  "two",
		Params: params,
	}
	events.Fire(payload)
	return nil, nil
}
