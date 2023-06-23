package lib

import (
	"encoding/json"
	"fmt"
)

type ProjectAPIResponse struct {
	Size     int    `json:"size"`
	Page     int    `json:"page"`
	Pagelen  int    `json:"pagelen"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Values   []struct {
		Type  string `json:"type"`
		Links struct {
			HTML struct {
				Href string `json:"href"`
				Name string `json:"name"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
				Name string `json:"name"`
			} `json:"avatar"`
		} `json:"links"`
		UUID  string `json:"uuid"`
		Key   string `json:"key"`
		Owner struct {
			Type string `json:"type"`
		} `json:"owner"`
		Name                    string `json:"name"`
		Description             string `json:"description"`
		IsPrivate               bool   `json:"is_private"`
		CreatedOn               string `json:"created_on"`
		UpdatedOn               string `json:"updated_on"`
		HasPubliclyVisibleRepos bool   `json:"has_publicly_visible_repos"`
	} `json:"values"`
}

type Project struct {
	Type  string `json:"type"`
	Links struct {
		HTML struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"avatar"`
	} `json:"links"`
	UUID  string `json:"uuid"`
	Key   string `json:"key"`
	Owner struct {
		Type string `json:"type"`
	} `json:"owner"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	IsPrivate               bool   `json:"is_private"`
	CreatedOn               string `json:"created_on"`
	UpdatedOn               string `json:"updated_on"`
	HasPubliclyVisibleRepos bool   `json:"has_publicly_visible_repos"`
}

type GetProjectsOutput struct {
	Projects []Project `json:"values"`
}

func (c *Client) GetProjects() ([]Project, error) {

	var (
		projects   = []Project{}
		pageLength = 50
		curPage    = 1
	)

	url := fmt.Sprintf(`https://api.bitbucket.org/2.0/workspaces/%s/projects`, c.workspace)

	for {
		res, err := c.request(url, pageLength, curPage)

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("Expected statuscode 200 but got: %d", res.StatusCode)
		}
		if err != nil {
			return nil, fmt.Errorf("Error retrieving projects: %s", err.Error())
		}

		defer res.Body.Close()

		output := &GetProjectsOutput{}
		err = json.NewDecoder(res.Body).Decode(output)
		if err != nil {
			return nil, fmt.Errorf("Error decoding response output: %s", err.Error())
		}

		projects = append(projects, output.Projects...)
		curPage++

		if len(output.Projects) < pageLength {
			break
		}
	}

	return projects, nil
}
