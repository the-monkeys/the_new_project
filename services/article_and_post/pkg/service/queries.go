package service

const (
	// getLast100Articles basically picks recent 100 published articles skipping the drafts
	getLast100Articles = ``

	// getArticlesByTags picks articles based on the tag name, latest first
	getArticlesByTags = `{
		"size": 100,
		"query": {
			"match": {
				"is_draft": "false"
			}
		},
		"_source": {
			"includes": [
				"id",
				"title",
				"author",
				"create_time",
				"quick_read"
			]
		}
	}`
)
