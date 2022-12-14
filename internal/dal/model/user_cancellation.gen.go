// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserCancellation = "user_cancellation"

// UserCancellation mapped from table <user_cancellation>
type UserCancellation struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID           string         `gorm:"column:uid;not null;default:0" json:"uid"`                     // 用户ID
	Reason        string         `gorm:"column:reason" json:"reason"`                                  // 申请理由
	ApplyTime     time.Time      `gorm:"column:apply_time;not null" json:"applyTime"`                  // 申请时间
	ConfirmTime   time.Time      `gorm:"column:confirm_time" json:"confirmTime"`                       // 确认时间
	ConfirmRemark string         `gorm:"column:confirm_remark" json:"confirmRemark"`                   // 确认备注
	Status        int32          `gorm:"column:status" json:"status"`                                  // 处理状态（1待处理，2注销通过，3注销驳回）
	CreatedAt     time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                           // 删除时间
}

// TableName UserCancellation's table name
func (*UserCancellation) TableName() string {
	return TableNameUserCancellation
}
