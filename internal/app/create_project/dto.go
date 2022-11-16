package create_project

type EventData struct {
	Token       string `json:"token"`
	ProjectName string `json:"project_name"`
}

type ServiceInput struct {
	Name        string
	Description string
	Owner       string
	Boilerplate string
	PipelineUrl string
}

type ServiceOutput struct {
	Url   string
	Error error
}

type RequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Boilerplate string `json:"boilerplate"`
}
