package article

type CreateArticleRequestBody struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Author      string   `json:"author"`
	AuthorEmail string   `json:"author_email"`
	IsDraft     bool     `json:"is_draft"`
	Tags        []string `json:"tags"`
}

type EditArticleRequestBody struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}
