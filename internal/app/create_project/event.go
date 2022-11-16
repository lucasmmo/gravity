package create_project

type Event struct {
	key  string
	data []byte
}

func (evt *Event) GetKey() string {
	return evt.key
}

func (evt *Event) GetData() interface{} {
	return evt.data
}
