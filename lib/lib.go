package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
		
)

type GetRepositoriesOutput struct {
	Repositories []Repository `json:"values"`
}

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

func GetProjects(workspace string, appPass string, appUser string) ([]Project, error) {
	projects := []Project{}
	pageLength := 50
	curPage := 1

	for {
		client := http.Client{Timeout: 5 * time.Second}

		url := fmt.Sprintf(`https://api.bitbucket.org/2.0/projects/%s`, workspace)
		req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
		if err != nil {
			return nil, fmt.Errorf("Error creating projects requests url: %s", err.Error())
		}

		q := req.URL.Query()
		q.Set("pagelen", strconv.Itoa(pageLength))
		q.Set("page", strconv.Itoa(curPage))
		req.URL.RawQuery = q.Encode()
		
		req.SetBasicAuth(appUser, appPass)

		res, err := client.Do(req)

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("Expected statuscode 200 but got: %d", res.StatusCode)
		}
		if err != nil {
			return nil, fmt.Errorf("Error retrieving projects: %s", err.Error())
		}

		defer res.Body.Close()

		getProjectsOutput := &GetProjectsOutput{}
		err = json.NewDecoder(res.Body).Decode(getProjectsOutput)
		if err != nil {
			return nil, fmt.Errorf("Error decoding response output: %s", err.Error())
		}

		projects = append(projects, getProjectsOutput.Projects...)
		curPage++

		if len(getProjectsOutput.Projects) < pageLength {
			break
		}
	}
	return projects, nil
}

type GetUsersOutput struct {
	Size     int    `json:"size"`
	Page     int    `json:"page"`
	Pagelen  int    `json:"pagelen"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Values   []struct {
		Type  string `json:"type"`
		Links struct {
			Self struct {
				Href string `json:"href"`
				Name string `json:"name"`
			} `json:"self"`
		} `json:"links"`
		User struct {
			Type string `json:"type"`
		} `json:"user"`
		Workspace struct {
			Type string `json:"type"`
		} `json:"workspace"`
	} `json:"values"`
}

type Project struct {
	Values []struct {
		Name        string `json:"name"`
		Key         string `json:"key"`
		ID          int    `json:"id"`
		Type        string `json:"type"`
		Public      bool   `json:"public"`
		Scope       string `json:"scope"`
		Description string `json:"description"`
		Namespace   string `json:"namespace"`
		Avatar      string `json:"avatar"`
	} `json:"values"`
	Size          int  `json:"size"`
	Limit         int  `json:"limit"`
	Start         int  `json:"start"`
	IsLastPage    bool `json:"isLastPage"`
	NextPageStart int  `json:"nextPageStart"`
}

type GetProjectsOutput struct {
	Projects []Project `json:"values"`
}

func GetProjects(workspace string, appPass string, appUser string) ([]Project, error) {
	repos := []Repository{}
	pageLength := 50
	curPage := 1

	for {
		client := http.Client{Timeout: 5 * time.Second}

		url := fmt.Sprintf(`https://api.bitbucket.org/2.0/repositories/%s`, workspace)
		req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
		if err != nil {
			return nil, fmt.Errorf("Error creating repositories requests url: %s", err.Error())
		}

		q := req.URL.Query()
		q.Set("pagelen", strconv.Itoa(pageLength))
		q.Set("page", strconv.Itoa(curPage))
		req.URL.RawQuery = q.Encode()
		
		req.SetBasicAuth(appUser, appPass)

		res, err := client.Do(req)

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("Expected statuscode 200 but got: %d", res.StatusCode)
		}
		if err != nil {
			return nil, fmt.Errorf("Error retrieving repositories: %s", err.Error())
		}

		defer res.Body.Close()

		getRepositoriesOutput := &GetRepositoriesOutput{}
		err = json.NewDecoder(res.Body).Decode(getRepositoriesOutput)
		if err != nil {
			return nil, fmt.Errorf("Error decoding response output: %s", err.Error())
		}

		repos = append(repos, getRepositoriesOutput.Repositories...)
		curPage++

		if len(getRepositoriesOutput.Repositories) < pageLength {
			break
		}
	}
	return repos, nil
}