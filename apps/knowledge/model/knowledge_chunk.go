package model

import "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"

type KnowledgeChunk struct {
	Base

	KnowledgeID    int64                        `gorm:"column:knowledge_id" json:"knowledge_id"`
	DocumentID     int64                        `gorm:"column:document_id" json:"document_id"`
	Content        string                       `gorm:"column:content" json:"content"`
	Order          int64                        `gorm:"column:order" json:"order"`
	RetrievalCount int64                        `gorm:"column:retrieval_count" json:"retrieval_count"`
	Status         knowledgesvc.EmbeddingStatus `gorm:"column:status" json:"status"`
}

type KnowledgeChunkOptional struct {
}
