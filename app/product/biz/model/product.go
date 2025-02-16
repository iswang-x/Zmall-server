package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture" gorm:"column:picture;comment:;size:64;"`
	Price       float32 `json:"price"`
	//商品和商品分类：多对多关系,通过一个product_category表实现
	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

func (P Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// 获取商品信息
func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return
}

// 搜索商品列表
func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?",
		"%"+q+"%", "%"+q+"%",
	).Error
	return
}

func NewPorductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}
