package model

import (
	"context"
	"gorm.io/gorm"
)

type ContextModel struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Timestamp int64
	Date      string `json:"date"`
	Msg       []byte `json:"private_key"`
}

type DateModel struct {
	db *gorm.DB
}

func NewDataModel(db *gorm.DB) *DateModel {
	return &DateModel{
		db: db,
	}
}

// Add 新增一条数据
func (dm *DateModel) Add(ctx context.Context, ct *ContextModel) error {
	return dm.db.WithContext(ctx).Create(ct).Error
}

// GetDate 获取date数量
func (dm *DateModel) GetDate(ctx context.Context) ([]*ContextModel, error) {
	var cm []*ContextModel
	err := dm.db.WithContext(ctx).Group("date").Find(&cm).Error
	return cm, err
}

// GetContextByDate 通过date获取数据
func (dm *DateModel) GetContextByDate(ctx context.Context, date string) ([]*ContextModel, error) {
	var cm []*ContextModel
	err := dm.db.WithContext(ctx).Where("date", date).Order("timestamp").Find(&cm).Error
	return cm, err
}
