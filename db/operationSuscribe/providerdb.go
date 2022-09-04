package operationSuscribe

import "github.com/sebarray/wiselink/model"

type SuscribeSql struct{}

type ISuscribe interface {
	CreateSuscribe(userId string, event model.Event) error
	ReadSuscribe(filter, id string) ([]model.Event, error)
}

var providers = map[string]ISuscribe{}

func init() {
	providers["mysql"] = SuscribeSql{}
}

func GetProvider(typeDb string) ISuscribe {
	return providers[typeDb]
}
