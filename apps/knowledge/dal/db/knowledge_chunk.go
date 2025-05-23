package db

import (
	"context"
	"math"

	"github.com/aiagt/aiagt/apps/knowledge/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/pkg/snowflake"
	"github.com/pkg/errors"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type KnowledgeChunkDao struct {
	m *model.KnowledgeChunk
}

// NewKnowledgeChunkDao make KnowledgeChunk dao
func NewKnowledgeChunkDao() *KnowledgeChunkDao {
	return &KnowledgeChunkDao{m: new(model.KnowledgeChunk)}
}

func (d *KnowledgeChunkDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get knowledge_chunk by id
func (d *KnowledgeChunkDao) GetByID(ctx context.Context, id int64) (*model.KnowledgeChunk, error) {
	var result model.KnowledgeChunk

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge_chunk dao get by id error")
	}

	return &result, nil
}

// GetByIDs get knowledge_chunk list by ids
func (d *KnowledgeChunkDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.KnowledgeChunk, error) {
	var result []*model.KnowledgeChunk

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge_chunk dao get by ids error")
	}

	return result, nil
}

// List get knowledge_chunk list
func (d *KnowledgeChunkDao) List(ctx context.Context, page *base.PaginationReq) ([]*model.KnowledgeChunk, *base.PaginationResp, error) {
	var (
		list   []*model.KnowledgeChunk
		total  int64
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "knowledge_chunk dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a knowledge_chunk record
func (d *KnowledgeChunkDao) Create(ctx context.Context, m *model.KnowledgeChunk) error {
	m.ID = snowflake.Generate().Int64()

	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_chunk dao create error")
	}

	return nil
}

// Update knowledge_chunk by id
func (d *KnowledgeChunkDao) Update(ctx context.Context, id int64, m *model.KnowledgeChunkOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_chunk dao update error")
	}

	return nil
}

// Delete delete knowledge_chunk by id
func (d *KnowledgeChunkDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_chunk dao delete error")
	}

	return nil
}

// DeleteByDocumentID delete knowledge_chunk by document_id
func (d *KnowledgeChunkDao) DeleteByDocumentID(ctx context.Context, documentID int64) error {
	err := d.db(ctx).Model(d.m).Where("document_id = ?", documentID).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_chunk dao delete by document id error")
	}

	return nil
}

// GetBy get by condition
func (d *KnowledgeChunkDao) GetBy(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (*model.KnowledgeChunk, error) {
	var result model.KnowledgeChunk

	err := d.db(ctx).Model(d.m).Scopes(scopes...).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge_chunk dao get by error")
	}

	return &result, nil
}

// ListBy list by condition
func (d *KnowledgeChunkDao) ListBy(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.KnowledgeChunk, error) {
	var result []*model.KnowledgeChunk

	err := d.db(ctx).Model(d.m).Scopes(scopes...).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge_chunk dao list by error")
	}

	return result, nil
}

func (d *KnowledgeChunkDao) CreateBatch(ctx context.Context, ms []*model.KnowledgeChunk) error {
	for _, m := range ms {
		m.ID = snowflake.Generate().Int64()
	}

	err := d.db(ctx).Model(d.m).CreateInBatches(ms, 100).Error
	if err != nil {
		return errors.Wrap(err, "knowledge_chunk dao create batch error")
	}

	return nil
}
