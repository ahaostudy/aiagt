package db

import (
	"context"
	"fmt"
	"github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
	"math"

	"github.com/aiagt/aiagt/apps/knowledge/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/pkg/snowflake"
	"github.com/pkg/errors"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type KnowledgeDao struct {
	m *model.Knowledge
}

// NewKnowledgeDao make Knowledge dao
func NewKnowledgeDao() *KnowledgeDao {
	return &KnowledgeDao{m: new(model.Knowledge)}
}

func (d *KnowledgeDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get knowledge by id
func (d *KnowledgeDao) GetByID(ctx context.Context, id int64) (*model.Knowledge, error) {
	var result model.Knowledge

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge dao get by id error")
	}

	return &result, nil
}

// GetByIDs get knowledge list by ids
func (d *KnowledgeDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.Knowledge, error) {
	var result []*model.Knowledge

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge dao get by ids error")
	}

	return result, nil
}

// List get knowledge list
func (d *KnowledgeDao) List(ctx context.Context, req *knowledgesvc.ListKnowledgeReq, userID int64) ([]*model.Knowledge, *base.PaginationResp, error) {
	var (
		list   []*model.Knowledge
		total  int64
		offset = int((req.Pagination.Page - 1) * req.Pagination.PageSize)
		limit  = int(req.Pagination.PageSize)
	)

	err := d.db(ctx).Model(d.m).Scopes(func(db *gorm.DB) *gorm.DB {
		if req.AuthorId == nil {
			db = db.Where("author_id = ? OR is_private = ?", userID, false)
		} else {
			if *req.AuthorId != userID {
				db = db.Where("author_id = ? AND is_private = ?", *req.AuthorId, false)
			} else {
				db = db.Where("author_id = ?", *req.AuthorId)
			}
		}
		if req.Name != nil {
			db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *req.Name))
		}
		return db
	}).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "knowledge dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(req.Pagination.PageSize)))
	pageResp := &base.PaginationResp{Page: req.Pagination.Page, PageSize: req.Pagination.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a knowledge record
func (d *KnowledgeDao) Create(ctx context.Context, m *model.Knowledge) error {
	m.ID = snowflake.Generate().Int64()

	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge dao create error")
	}

	return nil
}

// Update knowledge by id
func (d *KnowledgeDao) Update(ctx context.Context, id int64, m *model.KnowledgeOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge dao update error")
	}

	return nil
}

// Delete delete knowledge by id
func (d *KnowledgeDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge dao delete error")
	}

	return nil
}

// GetBy get by condition
func (d *KnowledgeDao) GetBy(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (*model.Knowledge, error) {
	var result model.Knowledge

	err := d.db(ctx).Model(d.m).Scopes(scopes...).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge dao get by error")
	}

	return &result, nil
}

// ListBy list by condition
func (d *KnowledgeDao) ListBy(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.Knowledge, error) {
	var result []*model.Knowledge

	err := d.db(ctx).Model(d.m).Scopes(scopes...).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "knowledge dao list by error")
	}

	return result, nil
}

func (d *KnowledgeDao) Save(ctx context.Context, m *model.Knowledge) error {
	if m.ID == 0 {
		return d.Create(ctx, m)
	}

	err := d.db(ctx).Save(m).Error
	if err != nil {
		return errors.Wrap(err, "knowledge dao save error")
	}

	return nil
}
