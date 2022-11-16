package git_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GetUserdata interface {
	GetUserdata(token string) (map[string]string, error)
}

type CreateRepository interface {
	CreateRepository(token, name, owner, description, boilerplate string) ([]string, error)
}

type CreateWebhook interface {
	CreateWebhook(token, owner, repoName, webhookUrl, webhookName string) error
}

type GithubClient struct {
	client *http.Client
}

func NewGithubClient(client *http.Client) *GithubClient {
	return &GithubClient{client}
}

func (cli *GithubClient) CreateRepository(token, name, description, owner, boilerplate string) ([]string, error) {
	reqBody, err := json.Marshal(CreateRepositoryByTemplateRequestBody{name, description})
	if err != nil {
		return []string{}, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.github.com/repos/%s/%s/generate", owner, boilerplate), bytes.NewBuffer(reqBody))
	if err != nil {
		return []string{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := cli.client.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return []string{}, errors.New("cannot create repository in Github")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	var responseObject CreateRepositoryByTemplateResponse
	if err := json.Unmarshal(body, &responseObject); err != nil {
		return []string{}, err
	}

	return []string{responseObject.GitUrl, responseObject.SshUrl, responseObject.HtmlUrl}, nil
}

func (cli *GithubClient) GetUserdata(token string) (map[string]string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return map[string]string{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := cli.client.Do(req)
	if err != nil {
		return map[string]string{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return map[string]string{}, errors.New("cannot get user data in Github")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]string{}, err
	}

	var responseObject map[string]string

	json.Unmarshal(body, &responseObject)

	return responseObject, nil
}

func (cli *GithubClient) CreateWebhook(token, owner, repoName, webhookUrl, webhookName string) error {
	reqBody, err := json.Marshal(CreateWebhookInRepositoryRequestBody{
		Name: webhookName,
		Config: struct {
			URL         string `json:"url"`
			InsecureSsl string `json:"insecure_ssl"`
		}{
			URL:         webhookUrl,
			InsecureSsl: "1",
		},
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", owner, repoName), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := cli.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("cannot create webhook in Github repository")
	}

	return nil
}
