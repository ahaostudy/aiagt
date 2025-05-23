package handler

import (
	"context"
	"github.com/aiagt/aiagt/common/cos"
	"io"

	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
)

// PreviewChunks implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) PreviewChunks(ctx context.Context, req *knowledgesvc.PreviewChunksReq) (resp *knowledgesvc.PreviewChunksResp, err error) {
	content := req.GetContent()

	if req.Url != nil {
		fileResp, err := cos.Cli().Object.Get(ctx, req.GetUrl(), nil)
		if err != nil {
			return nil, bizPreviewChunks.NewErr(err).Log(ctx, "get file content by url error")
		}
		defer fileResp.Body.Close()

		content, err = io.ReadAll(fileResp.Body)
		if err != nil {
			return nil, bizPreviewChunks.NewErr(err).Log(ctx, "read body error")
		}
	}

	parsedDocuments, err := s.parseDocument(ctx, req.GetType(), content)
	if err != nil {
		return nil, bizPreviewChunks.NewErr(err).Log(ctx, "parse document error")
	}

	splitDocuments, err := s.splitDocument(ctx, parsedDocuments, req.GetChunkSize(), req.GetOverlapSize(), req.GetSeparators())
	if err != nil {
		return nil, bizPreviewChunks.NewErr(err).Log(ctx, "split document error")
	}

	resp = new(knowledgesvc.PreviewChunksResp)

	for i, document := range splitDocuments {
		resp.Chunks = append(resp.Chunks, &knowledgesvc.KnowledgeChunk{
			Content: document.Content,
			Vector:  nil,
			Order:   int64(i + 1),
		})
	}

	return
}
