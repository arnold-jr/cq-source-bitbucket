package lib

import (
	"encoding/json"
	"fmt"
	"time"
)

type Repository struct {
	Type     string `json:"type"`
	FullName string `json:"full_name"`
	Links    struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
		Pullrequests struct {
			Href string `json:"href"`
		} `json:"pullrequests"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Forks struct {
			Href string `json:"href"`
		} `json:"forks"`
		Watchers struct {
			Href string `json:"href"`
		} `json:"watchers"`
		Branches struct {
			Href string `json:"href"`
		} `json:"branches"`
		Tags struct {
			Href string `json:"href"`
		} `json:"tags"`
		Downloads struct {
			Href string `json:"href"`
		} `json:"downloads"`
		Source struct {
			Href string `json:"href"`
		} `json:"source"`
		Clone []struct {
			Name string `json:"name"`
			Href string `json:"href"`
		} `json:"clone"`
		Issues struct {
			Href string `json:"href"`
		} `json:"issues"`
		Hooks struct {
			Href string `json:"href"`
		} `json:"hooks"`
	} `json:"links"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Scm         string `json:"scm"`
	Website     any    `json:"website"`
	Owner       struct {
		DisplayName string `json:"display_name"`
		Links       struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"links"`
		Type     string `json:"type"`
		UUID     string `json:"uuid"`
		Username string `json:"username"`
	} `json:"owner"`
	Workspace struct {
		Type  string `json:"type"`
		UUID  string `json:"uuid"`
		Name  string `json:"name"`
		Slug  string `json:"slug"`
		Links struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"workspace"`
	IsPrivate bool `json:"is_private"`
	Project   struct {
		Type  string `json:"type"`
		Key   string `json:"key"`
		UUID  string `json:"uuid"`
		Name  string `json:"name"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
	} `json:"project"`
	ForkPolicy string    `json:"fork_policy"`
	CreatedOn  time.Time `json:"created_on"`
	UpdatedOn  time.Time `json:"updated_on"`
	Size       int       `json:"size"`
	Language   string    `json:"language"`
	HasIssues  bool      `json:"has_issues"`
	HasWiki    bool      `json:"has_wiki"`
	UUID       string    `json:"uuid"`
	Mainbranch struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"mainbranch"`
	OverrideSettings struct {
		DefaultMergeStrategy bool `json:"default_merge_strategy"`
		BranchingModel       bool `json:"branching_model"`
	} `json:"override_settings"`
}

type GetRepositoriesOutput struct {
	Repositories []Repository `json:"values"`
}

func (c *Client) GetRepositories() ([]Repository, error) {

	var (
		repositories = []Repository{}
		pageLength   = 50
		curPage      = 1
	)

	url := fmt.Sprintf(`https://api.bitbucket.org/2.0/repositories/%s`, c.workspace)

	for {
		res, err := c.request(url, pageLength, curPage)

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("Expected statuscode 200 but got: %d", res.StatusCode)
		}
		if err != nil {
			return nil, fmt.Errorf("Error retrieving repositories: %s", err.Error())
		}

		defer res.Body.Close()

		output := &GetRepositoriesOutput{}
		err = json.NewDecoder(res.Body).Decode(output)
		if err != nil {
			return nil, fmt.Errorf("Error decoding response output: %s", err.Error())
		}

		repositories = append(repositories, output.Repositories...)
		curPage++

		if len(output.Repositories) < pageLength {
			break
		}
	}

	return repositories, nil
}
