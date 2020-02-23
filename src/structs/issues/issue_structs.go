package issues

import "time"

//IssuesURL url to search issues
const IssuesURL = "https://api.github.com/search/issues"

//IssueSearchResult search result struct
type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

//Issue issue struct
type Issue struct {
	Number  int
	HTMLURL string `json:"html_url"`
	Title   string
	State   string
	User    *User
	Created time.Time `json:"created_at"`
	Body    string
}

//User user struct
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
