// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSensitiveWordCategory = "sensitive_word_category"

// SensitiveWordCategory mapped from table <sensitive_word_category>
type SensitiveWordCategory struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`                                      // 名称
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                           // 删除时间
}

// TableName SensitiveWordCategory's table name
func (*SensitiveWordCategory) TableName() string {
	return TableNameSensitiveWordCategory
}
