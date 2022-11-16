package git_client

type CreateRepositoryByTemplateResponse struct {
	SshUrl  string `json:"ssh_url"`
	GitUrl  string `json:"git_url"`
	HtmlUrl string `json:"html_url"`
}

type CreateRepositoryByTemplateRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetUserdataResponse struct {
	Name string `json:"login"`
}

type CreateWebhookInRepositoryRequestBody struct {
	Name   string `json:"name"`
	Config struct {
		URL         string `json:"url"`
		InsecureSsl string `json:"insecure_ssl"`
	} `json:"config"`
}
