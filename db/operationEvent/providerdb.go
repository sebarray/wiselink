package operationEvent

import (
	"github.com/sebarray/wiselink/model"
)

type EventSql struct{}

type IEventCrud interface {
	CreateEvent(event model.Event) (model.Event, error)
	ReadEvents(filter model.Filter) ([]model.Event, error)
	DeleteEvent(id string) error
	UpdateEvent(event model.Event) (model.Event, error)
}

var providers = map[string]IEventCrud{}

func init() {
	providers["mysql"] = EventSql{}
}

func GetProvider(typeDb string) IEventCrud {
	return providers[typeDb]
}
