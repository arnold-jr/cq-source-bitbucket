package lib

import (
	"encoding/json"
	"fmt"
)

type MemberHref struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

func (u MemberHref) ToUser() User {
	return User{}
}

type GetUsersOutput struct {
	Size     int          `json:"size"`
	Page     int          `json:"page"`
	Pagelen  int          `json:"pagelen"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Values   []MemberHref `json:"values"`
}

type User struct {
	Type  string `json:"type"`
	Links struct {
		Avatar struct {
		} `json:"avatar"`
	} `json:"links"`
	CreatedOn   string `json:"created_on"`
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	UUID        string `json:"uuid"`
}

func (c *Client) GetUsers() ([]User, error) {

	var (
		memberHrefs = []MemberHref{}
		users 		= []User{}
		pageLength  = 50
		curPage     = 1
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

		memberHrefs = append(memberHrefs, output.Values...)
		curPage++

		if len(output.Values) < pageLength {
			break
		}

	}

	for _, memberHref := range memberHrefs {
		url := memberHref.Links.Self.Href
		fmt.Println(url)
		res, err := c.request(url, 1, 1)

		output := User{}
		err = json.NewDecoder(res.Body).Decode(&output)
		if err != nil {
			return nil, fmt.Errorf("Error decoding response output: %s", err.Error())
		}
		// decode body into output
		users = append(users, output)
		curPage++

		fmt.Println(err, res)
	}
	return users, nil
	// for each in range []MemberRef

	// make a request using the HREF
	// req := c.request(membr.href.somthing)

	// convert the body, to a User

	// add the User to tthe []user

	// return users, nil

}
