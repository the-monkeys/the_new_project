package article

type CreateArticleRequestBody struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	IsDraft bool     `json:"is_draft"`
	Tags    []string `json:"tags"`
}
