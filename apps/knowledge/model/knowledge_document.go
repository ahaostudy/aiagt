package model

import (
	"github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
)

type KnowledgeDocument struct {
	Base

	KnowledgeID    int64                        `gorm:"column:knowledge_id" json:"knowledge_id"`
	Name           string                       `gorm:"column:name" json:"name"`
	Url            string                       `gorm:"column:url" json:"url"`
	Type           string                       `gorm:"column:type" json:"type"`
	Preview        string                       `gorm:"column:preview" json:"preview"`
	Size           int64                        `gorm:"column:size" json:"size"`
	RetrievalCount int64                        `gorm:"column:retrieval_count" json:"retrieval_count"`
	ChunkSize      int64                        `gorm:"column:chunk_size" json:"chunk_size"`
	OverlapSize    int64                        `gorm:"column:overlap_size" json:"overlap_size"`
	Separators     []string                     `gorm:"column:separators;serializer:json" json:"separators"`
	Status         knowledgesvc.EmbeddingStatus `gorm:"column:status" json:"status"`
	FailedReason   string                       `gorm:"column:failed_reason" json:"failed_reason"`
}

type KnowledgeDocumentOptional struct {
}
