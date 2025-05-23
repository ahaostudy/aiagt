package controller

import (
	"context"
	"encoding/json"
	"github.com/aiagt/aiagt/common/hertz/router"
	"github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
	knowledgeservice "github.com/aiagt/aiagt/kitex_gen/knowledgesvc/knowledgeservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"io"
	"path/filepath"
)

func RegisterRouter(r *route.RouterGroup, cli knowledgeservice.Client) {
	r = r.Group("/knowledge")

	router.POST(r, "", cli.SaveKnowledge)
	router.PUT(r, "/:id", cli.SaveKnowledge)
	router.DELETE(r, "/:id", cli.DeleteKnowledge)
	router.GET(r, "/:id", cli.GetKnowledgeByID)
	router.POST(r, "/list", cli.ListKnowledge)

	document := r.Group("/document")

	router.GET(document, "/get_by_knowledge", cli.GetDocumentByKnowledge)
	router.PUT(document, "/:id", cli.UpdateDocument)
	router.DELETE(document, "/:id", cli.DeleteDocument)

	router.POST(document, "/upload", cli.UploadDocument, func(ctx context.Context, c *app.RequestContext) (*knowledgesvc.UploadDocumentReq, error) {
		var req knowledgesvc.UploadDocumentReq

		if err := c.BindAndValidate(&req); err != nil {
			return nil, err
		}

		// parse file
		file, err := c.FormFile("content")
		if err != nil {
			return nil, err
		}

		fileBody, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer fileBody.Close()

		fileBodyBytes, err := io.ReadAll(fileBody)
		if err != nil {
			return nil, err
		}

		// parse separators
		if len(req.Separators) > 0 {
			var seps []string
			err = json.Unmarshal([]byte(req.Separators[0]), &seps)
			if err != nil {
				return nil, err
			}

			req.Separators = seps
		}

		req.Content = fileBodyBytes
		req.Name = file.Filename
		req.Type = filepath.Ext(file.Filename)

		return &req, nil
	})

	router.POST(document, "/preview_chunks", cli.PreviewChunks, func(ctx context.Context, c *app.RequestContext) (*knowledgesvc.PreviewChunksReq, error) {
		var req knowledgesvc.PreviewChunksReq

		if err := c.BindAndValidate(&req); err != nil {
			return nil, err
		}

		// parse file
		if req.Url == nil || len(*req.Url) == 0 {
			file, err := c.FormFile("content")
			if err != nil {
				return nil, err
			}

			fileBody, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer fileBody.Close()

			fileBodyBytes, err := io.ReadAll(fileBody)
			if err != nil {
				return nil, err
			}

			req.Content = fileBodyBytes
			req.Type = filepath.Ext(file.Filename)
		}

		// parse separators
		if len(req.Separators) > 0 {
			var seps []string
			err := json.Unmarshal([]byte(req.Separators[0]), &seps)
			if err != nil {
				return nil, err
			}

			req.Separators = seps
		}

		return &req, nil
	})
}
