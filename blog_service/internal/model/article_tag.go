package model

type ArticleTag struct{
	*Model
	TagID string `json:"tag_id"`
	ArticleID string `json:"article_id"`
}

func (a ArticleTag) TableName()string{
	return "blog_article_tag"
}
