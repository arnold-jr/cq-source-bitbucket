package lib

import (
	"encoding/json"
	"fmt"
)

type Project struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Public      bool   `json:"public"`
	Scope       string `json:"scope"`
	Description string `json:"description"`
	Namespace   string `json:"namespace"`
	Avatar      string `json:"avatar"`
}


type GetProjectsOutput struct {
	Projects []Project `json:"values"`
}

func (c *Client) GetProjects() ([]Project, error) {

	var (
		projects = []Project{}
		pageLength   = 50
		curPage      = 1
	)

	url := fmt.Sprintf(`https://api.bitbucket.org/2.0/workspace/%s/projects`, c.workspace)

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