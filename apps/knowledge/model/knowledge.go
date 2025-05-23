package model

import "time"

type Knowledge struct {
	Base

	Name           string     `gorm:"column:name;type:varchar(255)" json:"name"`
	Description    string     `gorm:"column:description;type:text" json:"description"`
	AuthorID       int64      `gorm:"column:author_id;type:bigint" json:"author_id"`
	IsPrivate      bool       `gorm:"column:is_private" json:"is_private"`
	Logo           string     `gorm:"column:logo" json:"logo"`
	EmbedModelID   int64      `gorm:"column:embed_model_id;type:bigint" json:"embed_model_id"`
	TopK           int64      `gorm:"column:top_k" json:"top_k"`
	ScoreThreshold float64    `gorm:"column:score_threshold" json:"score_threshold"`
	PublishedAt    *time.Time `gorm:"column:published_at" json:"published_at"`
}

type KnowledgeOptional struct {
}
