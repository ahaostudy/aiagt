package db

import (
	"context"
	"github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
	"math"

	"github.com/aiagt/aiagt/apps/knowledge/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/pkg/snowflake"
	"github.com/pkg/errors"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type KnowledgeDocumentDao struct {
	m *model.KnowledgeDocument
}

// NewKnowledgeDocumentDao make KnowledgeDocument dao
func NewKnowledgeDocumentDao() *KnowledgeDocumentDao {
	return &KnowledgeDocumentDao{m: new(model.KnowledgeDocument)}
}

func (d *KnowledgeDocumentDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get knowledge_document by id
func (d *KnowledgeDocumentDao) GetByID(ctx context.Context, id int64) (*model.KnowledgeDocument, error) {
	var result model.KnowledgeDocument

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge_document dao get by id error")
	}

	return &result, nil
}

// GetByIDs get knowledge_document list by ids
func (d *KnowledgeDocumentDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.KnowledgeDocument, error) {
	var result []*model.KnowledgeDocument

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge_document dao get by ids error")
	}

	return result, nil
}

// List get knowledge_document list
func (d *KnowledgeDocumentDao) List(ctx context.Context, page *base.PaginationReq) ([]*model.KnowledgeDocument, *base.PaginationResp, error) {
	var (
		list   []*model.KnowledgeDocument
		total  int64
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "knowledge_document dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a knowledge_document record
func (d *KnowledgeDocumentDao) Create(ctx context.Context, m *model.KnowledgeDocument) error {
	m.ID = snowflake.Generate().Int64()

	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_document dao create error")
	}

	return nil
}

// Update knowledge_document by id
func (d *KnowledgeDocumentDao) Update(ctx context.Context, id int64, m *model.KnowledgeDocumentOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_document dao update error")
	}

	return nil
}

// Delete delete knowledge_document by id
func (d *KnowledgeDocumentDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_document dao delete error")
	}

	return nil
}

// GetBy get by condition
func (d *KnowledgeDocumentDao) GetBy(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (*model.KnowledgeDocument, error) {
	var result model.KnowledgeDocument

	err := d.db(ctx).Model(d.m).Scopes(scopes...).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge_document dao get by error")
	}

	return &result, nil
}

// ListBy list by condition
func (d *KnowledgeDocumentDao) ListBy(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.KnowledgeDocument, error) {
	var result []*model.KnowledgeDocument

	err := d.db(ctx).Model(d.m).Scopes(scopes...).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge_document dao list by error")
	}

	return result, nil
}

func (d *KnowledgeDocumentDao) UpdateStatus(ctx context.Context, id int64, status knowledgesvc.EmbeddingStatus, failedReason string) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(map[string]any{
		"status":        status,
		"failed_reason": failedReason,
	}).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_document dao update status error")
	}

	return nil
}

// Save a knowledge_document record
func (d *KnowledgeDocumentDao) Save(ctx context.Context, m *model.KnowledgeDocument) error {
	err := d.db(ctx).Save(m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_document dao save error")
	}

	return nil
}
