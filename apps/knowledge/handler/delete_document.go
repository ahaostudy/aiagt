package handler

import (
	"context"
	"fmt"

	"github.com/aiagt/aiagt/apps/knowledge/conf"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeleteDocument implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) DeleteDocument(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	document, err := s.documentDao.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, bizDeleteDocument.NewErr(err).Log(ctx, "get document by id error")
	}

	knowledge, err := s.knowledgeDao.GetByID(ctx, document.KnowledgeID)
	if err != nil {
		return nil, bizDeleteDocument.NewErr(err).Log(ctx, "get knowledge by id error")
	}

	if ctxutil.Forbidden(ctx, knowledge.AuthorID) {
		return nil, bizDeleteDocument.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
	}

	err = s.milvusCli.Delete(ctx, conf.Conf().Milvus.Collection, "", fmt.Sprintf("document_id == '%d'", req.GetId()))
	if err != nil {
		return nil, bizDeleteDocument.NewErr(err).Log(ctx, "delete document vectors error")
	}

	err = s.chunkDao.DeleteByDocumentID(ctx, req.GetId())
	if err != nil {
		return nil, bizDeleteDocument.NewErr(err).Log(ctx, "delete document chunks error")
	}

	err = s.documentDao.Delete(ctx, req.GetId())
	if err != nil {
		return nil, bizDeleteDocument.NewErr(err).Log(ctx, "delete document error")
	}

	return
}
