package store_project

type EventData struct {
	Token       string `json:"token"`
	ProjectName string `json:"project_name"`
}

type ServiceInput struct {
	ProjectName string
}
