namespace go knowledgesvc

include './base.thrift'
include './user.thrift'

enum EmbeddingStatus {
    PENDING,
    SUCCESS,
    FAILED
}

struct Knowledge {
    1: required i64 id
    2: required string name
    3: required string description
    4: required i64 author_id
    5: optional user.User author
    6: required bool is_private
    7: required string logo
    8: required i64 embed_model_id
    9: required i64 top_k
    10: optional double score_threshold
    11: required base.Time created_at
    12: required base.Time updated_at
    13: required base.Time published_at
}

struct KnowledgeDocument {
    1: required i64 id
    2: required i64 knowledge_id
    3: required string name
    4: required string url
    5: required string type
    6: required string preview
    7: required i64 size
    8: required i64 retrieval_count
    9: required base.Time updated_at
    10: optional i64 chunk_size = 1024
    11: optional i64 overlap_size = 50
    12: optional list<string> separators
    13: required EmbeddingStatus status
    14: optional string failed_reason
}

struct KnowledgeChunk {
    1: required string id
    2: required i64 knowledge_id
    3: required i64 document_id
    4: required string content
    5: optional list<double> vector
    6: required i64 order
    7: required i64 retrieval_count
    8: required EmbeddingStatus status
}

struct EmbeddingModel {
    1: required i64 id
    2: required string name
    3: required string description
    4: required string model_key
    5: required i64 dim
}

struct RetrievalReq {
    1: required list<i64> knowledge_ids
    2: required string query
}

struct RetrievalResp {
    1: required list<KnowledgeChunk> chunks
}

struct SaveKnowledgeReq {
    1: optional i64 id (go.tag='path:"id"')
    2: required string name
    3: required string description
    4: required bool is_private
    5: required i64 embed_model_id
    6: required i64 top_k
    7: optional double score_threshold
    8: required string logo
}

struct UploadDocumentReq {
    1: required i64 knowledge_id (go.tag='form:"knowledge_id"')
    2: required string name (go.tag='form:"name"')
    3: required binary content (go.tag='form:"content"')
    4: required string type (go.tag='form:"type"')
    5: optional i64 chunk_size = 1024 (go.tag='form:"chunk_size"')
    6: optional i64 overlap_size = 50 (go.tag='form:"overlap_size"')
    7: optional list<string> separators (go.tag='form:"separators"')
}

struct UpdateDocumentReq {
    1: required i64 id (go.tag='path:"id"')
    2: optional i64 chunk_size = 1024
    3: optional i64 overlap_size = 50
    4: optional list<string> separators
}

struct GetDocumentByKnowledgeResp {
    1: required list<KnowledgeDocument> documents
}

struct ListKnowledgeReq {
    1: required base.PaginationReq pagination
    2: optional i64 author_id
    3: optional string name
}

struct ListKnowledgeResp {
    1: required base.PaginationResp pagination
    2: required list<Knowledge> knowledge_list
}

struct PreviewChunksReq {
    1: optional binary content (go.tag='form:"content"')
    2: optional string url (go.tag='form:"url"')
    3: optional i64 chunk_size = 1024 (go.tag='form:"chunk_size"')
    4: optional i64 overlap_size = 50 (go.tag='form:"overlap_size"')
    5: optional list<string> separators (go.tag='form:"separators"')
    6: required string type (go.tag='form:"type"')
}

struct PreviewChunksResp {
    1: required list<KnowledgeChunk> chunks
}

service KnowledgeService {
    RetrievalResp Retrieval(1: RetrievalReq req)
    ListKnowledgeResp ListKnowledge(1: ListKnowledgeReq req)
    Knowledge GetKnowledgeByID(1: base.IDReq req)
    list<Knowledge> GetKnowledgeByIDs(1: base.IDsReq req)
    base.Empty SaveKnowledge(1: SaveKnowledgeReq req)
    base.Empty DeleteKnowledge(1: base.IDReq req)
    GetDocumentByKnowledgeResp GetDocumentByKnowledge(1: base.IDReq req)
    base.Empty UploadDocument(1: UploadDocumentReq req)
    base.Empty UpdateDocument(1: UpdateDocumentReq req)
    base.Empty DeleteDocument(1: base.IDReq req)
    PreviewChunksResp PreviewChunks(1: PreviewChunksReq req)
}