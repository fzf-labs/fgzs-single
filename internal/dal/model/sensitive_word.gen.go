// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSensitiveWord = "sensitive_word"

// SensitiveWord mapped from table <sensitive_word>
type SensitiveWord struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CategoryID int64          `gorm:"column:category_id" json:"categoryId"`                         // 分类ID
	Text       string         `gorm:"column:text" json:"text"`                                      // 敏感词
	CreatedAt  time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                           // 删除时间
}

// TableName SensitiveWord's table name
func (*SensitiveWord) TableName() string {
	return TableNameSensitiveWord
}
