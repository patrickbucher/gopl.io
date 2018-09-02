package ex14

type SearchResult struct {
	SearchTerms string  // not in JSON structure
	Count       int     `json:"total_count"`
	Issues      []Issue `json:"items"`
}

type Issue struct {
	Id            int       `json:"id"`
	Number        int       `json:"number"`
	RepositoryURL string    `json:"repository_url"`
	Title         string    `json:"title"`
	User          User      `json:"user"`
	Labels        []Label   `json:"labels"`
	State         string    `json:"state"`
	Milestone     Milestone `json:"milestone"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
	ClosedAt      string    `json:"closed_at"`
	Body          string    `json:"body"`
}

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"login"`
	ProfileURL string `json:"url"`
}

type Label struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Milestone struct {
	Id          int    `json:"id"`
	Number      int    `json:"number"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
