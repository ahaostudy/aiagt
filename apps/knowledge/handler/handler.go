package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/knowledge/dal/db"
	"github.com/aiagt/aiagt/apps/knowledge/pkg/rag"
	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	milvusretriever "github.com/cloudwego/eino-ext/components/retriever/milvus"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/pkg/errors"
)

// KnowledgeServiceImpl implements the last service interface defined in the IDL.
type KnowledgeServiceImpl struct {
	knowledgeDao *db.KnowledgeDao
	documentDao  *db.KnowledgeDocumentDao
	chunkDao     *db.KnowledgeChunkDao
	milvusCli    client.Client
}

func NewKnowledgeServiceImpl(knowledgeDao *db.KnowledgeDao, documentDao *db.KnowledgeDocumentDao, chunkDao *db.KnowledgeChunkDao) *KnowledgeServiceImpl {
	initServiceBusiness(7)

	return &KnowledgeServiceImpl{
		knowledgeDao: knowledgeDao,
		documentDao:  documentDao,
		chunkDao:     chunkDao,
	}
}

func (s *KnowledgeServiceImpl) SetMilvusClient(milvusClient client.Client) {
	s.milvusCli = milvusClient
}

func (s *KnowledgeServiceImpl) indexer(ctx context.Context) (*milvus.Indexer, error) {
	indexer, err := rag.NewIndexer(ctx, s.milvusCli)
	if err != nil {
		return nil, errors.Wrap(err, "new indexer error")
	}

	return indexer, nil
}

func (s *KnowledgeServiceImpl) retriever(ctx context.Context, topK int, scoreThreshold float64) (*milvusretriever.Retriever, error) {
	retriever, err := rag.NewRetriever(ctx, s.milvusCli, topK, scoreThreshold)
	if err != nil {
		return nil, errors.Wrap(err, "new retriever error")
	}

	return retriever, nil
}
