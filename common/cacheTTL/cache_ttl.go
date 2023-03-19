package cacheTTL

import "time"

var (
	MAXTTL time.Duration = time.Hour * 24 * 365
)

const (
	ArticleCommentByAid = 3600 * 24
	ArticleCommentByCid = 3600 * 24
)
