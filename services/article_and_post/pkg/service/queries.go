package service

const (

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

// getLast100Articles basically picks recent 100 published articles skipping the drafts
func getLast100Articles() string {
	return `{
		"size": 100,
		"sort": {
			"create_time": {
				"order": "desc"
			}
		},
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
				"quick_read",
				"author_email"
			]
		}
	}`
}
