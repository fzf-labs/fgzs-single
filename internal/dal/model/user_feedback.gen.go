// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserFeedback = "user_feedback"

// UserFeedback mapped from table <user_feedback>
type UserFeedback struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID       string         `gorm:"column:uid" json:"uid"`                                        // 用户ID
	Content   string         `gorm:"column:content" json:"content"`                                // 反馈内容
	Remark    string         `gorm:"column:remark" json:"remark"`                                  // 备注内容
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                           // 删除时间
}

// TableName UserFeedback's table name
func (*UserFeedback) TableName() string {
	return TableNameUserFeedback
}
