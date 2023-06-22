package lib

type GetProjectsOutput struct {
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
