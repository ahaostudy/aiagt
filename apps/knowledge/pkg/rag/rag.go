package rag

import (
	"context"
	"encoding/binary"
	"github.com/aiagt/aiagt/apps/knowledge/conf"
	"github.com/cloudwego/eino-ext/components/embedding/openai"
	milvusindexer "github.com/cloudwego/eino-ext/components/indexer/milvus"
	milvusretriever "github.com/cloudwego/eino-ext/components/retriever/milvus"
	"github.com/cloudwego/eino/schema"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	entitysdk "github.com/milvus-io/milvus-sdk-go/v2/entity"
	entitycli "github.com/milvus-io/milvus/client/v2/entity"
	"github.com/pkg/errors"
	"math"
	"time"
)

func NewMilvusClient() (client.Client, error) {
	cfg := conf.Conf().Milvus

	cli, err := client.NewClient(context.TODO(), client.Config{
		Address:  cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
	})
	if err != nil {
		return nil, errors.Wrap(err, "new milvus client error")
	}

	return cli, nil
}

func NewIndexer(ctx context.Context, cli client.Client) (*milvusindexer.Indexer, error) {
	cfg := conf.Conf().Embedding

	embedder, err := openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
		APIKey:  cfg.APIKey,
		Model:   cfg.Model,
		BaseURL: cfg.BaseURL,
		Timeout: timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "new embedder error")
	}

	indexer, err := milvusindexer.NewIndexer(ctx, &milvusindexer.IndexerConfig{
		Client:            cli,
		Embedding:         embedder,
		Collection:        collection,
		Fields:            KnowledgeFieldsForSDK(),
		DocumentConverter: IndexerDocumentConverter(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "new milvus indexer error")
	}

	return indexer, nil
}

func NewRetriever(ctx context.Context, cli client.Client, topK int, scoreThreshold float64) (*milvusretriever.Retriever, error) {
	cfg := conf.Conf().Embedding

	embdder, err := openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
		APIKey:  cfg.APIKey,
		Model:   cfg.Model,
		BaseURL: cfg.BaseURL,
		Timeout: timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "new embedder error")
	}

	retriever, err := milvusretriever.NewRetriever(ctx, &milvusretriever.RetrieverConfig{
		Client:         cli,
		Collection:     collection,
		VectorField:    vectorField,
		OutputFields:   outputFields,
		TopK:           topK,
		ScoreThreshold: scoreThreshold,
		Embedding:      embdder,
	})
	if err != nil {
		return nil, errors.Wrap(err, "new milvus retriever error")
	}

	return retriever, nil
}

var (
	collection   = "knowledge"
	vectorField  = "vector"
	outputFields = []string{"id", "knowledge_id", "document_id", "content", "vector"}
	timeout      = time.Minute
)

type KnowledgeSchame struct {
	ID          string `json:"id" milvus:"name:id"`
	KnowledgeID string `json:"knowledge_id" milvus:"name:knowledge_id"`
	DocumentID  string `json:"document_id" milvus:"name:document_id"`
	Content     string `json:"content" milvus:"name:content"`
	Vector      []byte `json:"vector" milvus:"name:vector"`
}

func KnowledgeFieldsForSDK() []*entitysdk.Field {
	return []*entitysdk.Field{
		entitysdk.NewField().
			WithName("id").
			WithIsPrimaryKey(true).
			WithDataType(entitysdk.FieldTypeVarChar).
			WithMaxLength(255),
		entitysdk.NewField().
			WithName("knowledge_id").
			WithIsPrimaryKey(false).
			WithDataType(entitysdk.FieldTypeVarChar).
			WithMaxLength(255),
		entitysdk.NewField().
			WithName("document_id").
			WithIsPrimaryKey(false).
			WithDataType(entitysdk.FieldTypeVarChar).
			WithMaxLength(255),
		entitysdk.NewField().
			WithName("content").
			WithIsPrimaryKey(false).
			WithDataType(entitysdk.FieldTypeVarChar).
			WithMaxLength(4096),
		entitysdk.NewField().
			WithName("vector").
			WithIsPrimaryKey(false).
			WithDataType(entitysdk.FieldTypeBinaryVector).
			WithDim(49152),
	}
}

func KnowledgeFieldsForCli() []*entitycli.Field {
	return []*entitycli.Field{
		entitycli.NewField().
			WithName("id").
			WithIsPrimaryKey(true).
			WithDataType(entitycli.FieldTypeVarChar).
			WithMaxLength(255),
		entitycli.NewField().
			WithName("knowledge_id").
			WithIsPrimaryKey(false).
			WithDataType(entitycli.FieldTypeVarChar).
			WithMaxLength(255),
		entitycli.NewField().
			WithName("document_id").
			WithIsPrimaryKey(false).
			WithDataType(entitycli.FieldTypeVarChar).
			WithMaxLength(255),
		entitycli.NewField().
			WithName("content").
			WithIsPrimaryKey(false).
			WithDataType(entitycli.FieldTypeVarChar).
			WithMaxLength(4096),
		entitycli.NewField().
			WithName("vector").
			WithIsPrimaryKey(false).
			WithDataType(entitycli.FieldTypeBinaryVector).
			WithDim(49152),
	}
}

func IndexerDocumentConverter() func(ctx context.Context, docs []*schema.Document, vectors [][]float64) ([]interface{}, error) {
	return func(ctx context.Context, docs []*schema.Document, vectors [][]float64) ([]interface{}, error) {
		em := make([]KnowledgeSchame, 0, len(docs))
		rows := make([]interface{}, 0, len(docs))

		for _, doc := range docs {
			knowledgeID, _ := doc.MetaData["knowledge_id"].(string)
			documentID, _ := doc.MetaData["document_id"].(string)
			em = append(em, KnowledgeSchame{
				ID:          doc.ID,
				KnowledgeID: knowledgeID,
				DocumentID:  documentID,
				Content:     doc.Content,
			})
		}

		for idx, vec := range vectors {
			em[idx].Vector = vector2Bytes(vec)
			rows = append(rows, &em[idx])
		}

		return rows, nil
	}
}

// vector2Bytes converts vector to bytes
func vector2Bytes(vector []float64) []byte {
	float32Arr := make([]float32, len(vector))
	for i, v := range vector {
		float32Arr[i] = float32(v)
	}
	bytes := make([]byte, len(float32Arr)*4)
	for i, v := range float32Arr {
		binary.LittleEndian.PutUint32(bytes[i*4:], math.Float32bits(v))
	}
	return bytes
}
