package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	db  *gorm.DB
	ctx context.Context
}

// 根据分类名称获取商品列表
func (c CategoryQuery) GetProductsByCategoryName(name string) (categories []*Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&categories).Error //Preload预加载分类包含的商品
	return
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		db:  db,
		ctx: ctx,
	}
}
