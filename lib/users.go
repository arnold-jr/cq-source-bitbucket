package lib

import (
	"encoding/json"
	"fmt"
)

type User struct {
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
}

type GetUsersOutput struct {
	Size     int    `json:"size"`
	Page     int    `json:"page"`
	Pagelen  int    `json:"pagelen"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Values   []User `json:"values"`
}

func (c *Client) GetUsers() ([]User, error) {

	var (
		users      = []User{}
		pageLength = 50
		curPage    = 1
	)

	url := fmt.Sprintf(`https://api.bitbucket.org/2.0/workspaces/%s/members`, c.workspace)

	for {
		res, err := c.request(url, pageLength, curPage)

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("Expected statuscode 200 but got: %d", res.StatusCode)
		}
		if err != nil {
			return nil, fmt.Errorf("Error retrieving users: %s", err.Error())
		}

		defer res.Body.Close()

		output := &GetUsersOutput{}
		err = json.NewDecoder(res.Body).Decode(output)
		if err != nil {
			return nil, fmt.Errorf("Error decoding response output: %s", err.Error())
		}

		users = append(users, output.Values...)
		curPage++

		if len(output.Values) < pageLength {
			break
		}
	
	}
	return users, nil
}