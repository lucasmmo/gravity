package store_project

import "encoding/json"

type listener struct {
	data    []byte
	service *service
}

func NewListener(service *service) *listener {
	return &listener{service: service}
}

func (l *listener) Handler() error {
	var input EventData

	if err := json.Unmarshal(l.data, &input); err != nil {
		return err
	}

	return l.service.Execute(input.Token, input.ProjectName)
}

func (l *listener) SetData(data interface{}) {
	l.data = data.([]byte)
}
