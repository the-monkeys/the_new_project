package blogsandposts

type CreatePostRequestBody struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Author    string   `json:"author"`
	AuthorId  string   `json:"author_id"`
	Published bool     `json:"published"`
	Tags      []string `json:"tags"`
}

type EditArticleRequestBody struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}
